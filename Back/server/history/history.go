package history

import (
    data "OSINT/Back/server/data"
    "OSINT/Back/server/logs"
    structure "OSINT/Back/server/structure"
    "net/http"
    "time"
    "go.uber.org/zap"
)

var logger = logs.GetLog(logs.GetLogConfig())

func GetHistory(w http.ResponseWriter, r *http.Request) {
    session, err := data.Store.Get(r, "data")
    if err != nil {
        logger.Error("Failed to get the session : ", zap.Error(err))
        return
    }
    _, ok := session.Values["username"].(string)
    if !ok {
        structure.TplData.ProcessMessage = "Vous n'êtes pas connecté !"
        http.Redirect(w, r, "/", http.StatusSeeOther)
        return
    }
    //Get the history of the user
    // history := data.GetHistory(session.Values["username"].(string))
    // structure.TplData.History = history
    if r.Method == "GET" {
        var research string
        var timestamp time.Time

        // err := data.Bd.QueryRow("SELECT research,timestamp FROM History WHERE userId = ?", username).Scan(&research, &timestamp)
        // err := data.Bd.QueryRow("SELECT research,timestamp FROM History WHERE userId = ?", username).Scan(&research, &timestamp)
        if err != nil {
            logger.Error("Failed to retrieve info user : ", zap.Error(err))
        }
        structure.TplData.History.Research = research
        structure.TplData.History.Timestamp = timestamp
        return
    } else {
        structure.TplData.ProcessMessage = "Your request is not valid ! : please retry "
        http.Redirect(w, r, "/", http.StatusSeeOther)
        return
    }
}