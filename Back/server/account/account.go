package account

import (
    data "OSINT/Back/server/data"
    logs "OSINT/Back/server/logs"
    structure "OSINT/Back/server/structure"
    "net/http"
    "time"

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
        http.Redirect(w, r, "/", http.StatusSeeOther)
        return
    }
    if r.Method == "GET" {

        var nom, prenom, email, icon string
        var age int

        // err := data.Bd.QueryRow("SELECT nom, prenom, email, age, icon FROM Utilisateurs WHERE username = ?", username).Scan(&nom, &prenom, &email, &age, &icon)

        if err != nil {
            logger.Error("Failed to retrieve info user : ", zap.Error(err))
        }
        rows, err := data.Bd.Query("SELECT research, timestamp FROM History WHERE username = ?", username)
        if err != nil {
            logger.Error("Failed to retrieve info history : ", zap.Error(err))
            return
        }
        defer rows.Close()

        var researches []string
        var timestamps []time.Time

        for rows.Next() {
            var research string
            var timestamp time.Time
            err := rows.Scan(&research, &timestamp)
            if err != nil {
                logger.Error("Failed to scan row: ", zap.Error(err))
                continue
            }
            researches = append(researches, research)
            timestamps = append(timestamps, timestamp)
        }

        if err := rows.Err(); err != nil {
            logger.Error("Error after iterating rows: ", zap.Error(err))
        }

        structure.TplData.History.Research = researches
        structure.TplData.History.Timestamp = timestamps
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
