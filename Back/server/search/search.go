package search

import (

    data "OSINT/Back/server/data"
    api "OSINT/Back/server/api"
    logs "OSINT/Back/server/logs"
    // structure "OSINT/Back/server/structure"
    "net/http"
    "go.uber.org/zap"
	"time"
)

var logger = logs.GetLog(logs.GetLogConfig())

func Search(w http.ResponseWriter, r *http.Request){
    logger.Info("tentativederecherche")
	session, err := data.Store.Get(r, "data")
    if err != nil {
        logger.Error("Failed to get the session : ", zap.Error(err))
        return
    }
    username, isUser := session.Values["username"].(string)
	if isUser{
        logger.Info("vous etes bien connecté")
		if r.Method == "POST" {
            logger.Info("Methode post succès")
            research := r.FormValue("search")
            api.Extract_Api("research")
            maintenant := time.Now()
            timestamp := maintenant.Format(time.RFC3339)
            _, errAddHistory := data.Bd.Exec("INSERT INTO History (username, research, timestamp) VALUES (?, ?, ?)", username, research, timestamp)

	    if errAddHistory != nil {
            // Gérer l'erreur, par exemple, en la journalisant ou en envoyant une réponse d'erreur au client
            logger.Info("Erreur lors de l'ajout à l'historique:", zap.Error(errAddHistory))
            http.Error(w, "Erreur lors de l'ajout à l'historique", http.StatusInternalServerError)
            return
        }

        // Si tout s'est bien passé, vous pouvez envoyer une réponse de succès au client
        w.WriteHeader(http.StatusOK)
        logger.Info("Recherche ajoutée à l'historique avec succès")
	}
	}
	
}