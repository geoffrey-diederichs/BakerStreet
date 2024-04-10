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
			now := time.Now()
			
			 last7days := now.Add(-7 * 24 * time.Hour)

            if err != nil {
                logger.Error("Failed to scan row: ", zap.Error(err))
                continue
            }
            researches = append(researches, research)
            timestamps = append(timestamps, last7days)
        }

        if err := rows.Err(); err != nil {
            logger.Error("Error after iterating rows: ", zap.Error(err))
        }

       var Yesterday []string
        var Last7days []string
        var thisMonth []string
        var Today []string

        today := time.Now().Truncate(24 * time.Hour)
        yesterday := today.Add(-24 * time.Hour)
        last7days := today.Add(-7 * 24 * time.Hour)
        firstOfMonth := time.Date(today.Year(), today.Month(), 1, 0, 0, 0, 0, today.Location())

        for i, ts := range timestamps {
            if ts.After(today) {
                Today = append(Today, researches[i])
            } else if ts.After(yesterday) {
                Yesterday = append(Yesterday, researches[i])
            } else if ts.After(last7days) {
                Last7days = append(Last7days, researches[i])
            } else if ts.After(firstOfMonth) {
                thisMonth = append(thisMonth, researches[i])
            }
        }

        // Envoyer les données au modèle
        structure.TplData.History.Today = Today
        structure.TplData.History.Yesterday = Yesterday
        structure.TplData.History.Last7Days = Last7days
        structure.TplData.History.ThisMonth = thisMonth

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
