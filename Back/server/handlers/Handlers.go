package handlers

import (
	account "OSINT/Back/server/account"
	auth "OSINT/Back/server/authentification"
	logs "OSINT/Back/server/logs"
	structure "OSINT/Back/server/structure"
	history "OSINT/Back/server/history"
	"html/template"
	"net/http"
	"go.uber.org/zap"
)

var tpl *template.Template

var logger = logs.GetLog(logs.GetLogConfig())

func init() {
	tpl = template.Must(template.ParseGlob("Front/pages/*.html"))
}

func WelcomeHandler(w http.ResponseWriter, r *http.Request) {
	errTpl := tpl.ExecuteTemplate(w, "accueil.html", structure.TplData)
	if errTpl != nil {
		logger.Error("", zap.Error(errTpl))
	}
}

func EnregistrementHandler(w http.ResponseWriter, r *http.Request) {
	auth.Enregistrement(w, r)
	errTpl := tpl.ExecuteTemplate(w, "register.html", structure.TplData)
	if errTpl != nil {
		logger.Error("", zap.Error(errTpl))
	}
}

// func PasswordModifyHandler(w http.ResponseWriter, r *http.Request) {
// 	auth.Enregistrement(w, r)
// 	fmt.Println("PasswordModify : " + structure.TplData.ProcessMessage)
// 	errTpl := tpl.ExecuteTemplate(w, "password_modify.html", structure.TplData)
// 	if errTpl != nil {
// 		fmt.Println(errTpl)
// 	}
// }

func LoginHandler(w http.ResponseWriter, r *http.Request) {
	auth.Login(w, r)
	errTpl := tpl.ExecuteTemplate(w, "login.html", structure.TplData)
	if errTpl != nil {
		logger.Error("", zap.Error(errTpl))
	}
}

func LogoutHandler(w http.ResponseWriter, r *http.Request) {
	auth.Logout(w, r)
	errTpl := tpl.ExecuteTemplate(w, "accueil.html", structure.TplData)
	if errTpl != nil {
		logger.Error("", zap.Error(errTpl))
	}
}

func SearchHandler(w http.ResponseWriter, r *http.Request) {
	// input := r.FormValue("input")
}

func HistoryHandler(w http.ResponseWriter, r *http.Request) {
	history.GetHistory(w, r)
	errTpl := tpl.ExecuteTemplate(w, "history.html", structure.TplData)
	if errTpl != nil {
		logger.Error("", zap.Error(errTpl))
	}
}

func AccountHandler(w http.ResponseWriter, r *http.Request) {
	account.GetUser(w, r)
	errTpl := tpl.ExecuteTemplate(w, "account.html", structure.TplData)
	if errTpl != nil {
		logger.Error("", zap.Error(errTpl))
	}
}
