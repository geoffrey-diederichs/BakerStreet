package search

import (
	data "OSINT/Back/server/data"
	logs "OSINT/Back/server/logs"
	"net/http"
	"time"
	"go.uber.org/zap"
)

var logger = logs.GetLog(logs.GetLogConfig())

func GetInput(w http.ResponseWriter, r *http.Request) string {
	session, err := data.Store.Get(r, "data")
	if err != nil {
		logger.Error("Failed to get the session : ", zap.Error(err))
		return ""
	}
	username, isUser := session.Values["username"].(string)
	if isUser {
		logger.Info("vous etes bien connecté en tant que %s",zap.String("",username))
		if r.Method == "POST" {
			research := r.FormValue("search")

			if research != "" {
				logger.Info("", zap.String("input found :", research))
			} else {
				logger.Info("", zap.String("input user is empty:", research))
				return ""
			}

			maintenant := time.Now()
			timestamp := maintenant.Format(time.RFC3339)
			_, errAddHistory := data.Bd.Exec("INSERT INTO History (username, research, timestamp) VALUES (?, ?, ?)", username, research, timestamp)

			if errAddHistory != nil {
				// Gérer l'erreur, par exemple, en la journalisant ou en envoyant une réponse d'erreur au client
				logger.Error("Erreur lors de l'ajout à l'historique:", zap.Error(errAddHistory))
				return ""
			}else{
                logger.Info("",zap.String("recherche",research))
                logger.Info("",zap.String("ajouté à l'historique utilisateur le",timestamp))
            }

			return research
		} else {
			logger.Error("mauvaise requete utilisateur")
		}
	}

	return ""
}
