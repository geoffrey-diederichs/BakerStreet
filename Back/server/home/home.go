package account

import (
    data "OSINT/Back/server/data"
    logs "OSINT/Back/server/logs"
    structure "OSINT/Back/server/structure"
    "net/http"

    "go.uber.org/zap"
)

var logger = logs.GetLog(logs.GetLogConfig())

func GetUser(w http.ResponseWriter, r *http.Request) {
    session, err := data.Store.Get(r, "data")
    if err != nil {
        logger.Error("Failed to get the session : ", zap.Error(err))
        return
    }
    username, isUser := session.Values["username"].(string)
    if !isUser {
        structure.TplData.ProcessMessage = "Vous n'êtes pas connecté !"
        return
    }
    if r.Method == "GET" {

        var nom, prenom, email, icon string
        var age int

        err := data.Bd.QueryRow("SELECT nom, prenom, email, age, icon FROM Utilisateurs WHERE username = ?", username).Scan(&nom, &prenom, &email, &age, &icon)

        if err != nil {
            logger.Error("Failed to retrieve info user : ", zap.Error(err))
        }

        structure.TplData.User.Username = username
        structure.TplData.User.Nom = nom
        structure.TplData.User.Prenom = prenom
        structure.TplData.User.Email = email
        structure.TplData.User.Age = age
        if icon != "" {
            structure.TplData.User.Icon = icon
        }
        return
    } else {
        structure.TplData.ProcessMessage = "Your request is not valid ! : please retry "
        http.Redirect(w, r, "/", http.StatusSeeOther)
        return
    }
}
