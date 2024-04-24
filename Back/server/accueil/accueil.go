package accueil

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
	if isUser {
		structure.TplData.ProcessMessage = "Bienvenue !"
		var icon string
		err := data.Bd.QueryRow("SELECT icon FROM Utilisateurs WHERE username = ?", username).Scan(&icon)
		if err != nil {
			logger.Error("Failed to retrieve info user : ", zap.Error(err))
		}
		structure.TplData.User.Username = username
		if icon != "" {
			structure.TplData.User.Icon = icon
		}
		return
	}
}
