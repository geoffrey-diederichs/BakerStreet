package handlers

import (
	account "OSINT/Back/server/account"
	home "OSINT/Back/server/home"
	auth "OSINT/Back/server/authentification"
	logs "OSINT/Back/server/logs"
	search "OSINT/Back/server/search"
	structure "OSINT/Back/server/structure"
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
	home.GetUser(w, r)
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
	search.Search(w, r)
	errTpl := tpl.ExecuteTemplate(w, "recherche.html", structure.TplData)
	if errTpl != nil {
		logger.Error("", zap.Error(errTpl))
	}
}

func AccountHandler(w http.ResponseWriter, r *http.Request) {
	account.GetUser(w, r)
	errTpl := tpl.ExecuteTemplate(w, "profil1.html", structure.TplData)
	if errTpl != nil {
		logger.Error("", zap.Error(errTpl))
	}
}
