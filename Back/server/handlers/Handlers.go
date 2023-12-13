package handlers

import (
	auth "OSINT/Back/server/authentification"
	structure "OSINT/Back/server/structure"
	"fmt"
	"html/template"
	"net/http"
)

var tpl *template.Template

func init() {
	tpl = template.Must(template.ParseGlob("Front/pages/*.html"))
}

func WelcomeHandler(w http.ResponseWriter, r *http.Request) {
	errTpl := tpl.ExecuteTemplate(w, "accueil.html", structure.TplData)
	if errTpl != nil {
		fmt.Println(errTpl)
	}
}

func EnregistrementHandler(w http.ResponseWriter, r *http.Request) {
	auth.Enregistrement(w, r)
	fmt.Println("Enregistrement : " + structure.TplData.ProcessMessage)
	errTpl := tpl.ExecuteTemplate(w, "register.html", structure.TplData)
	if errTpl != nil {
		fmt.Println(errTpl)
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
		fmt.Println(errTpl)
	}

}

func LogoutHandler(w http.ResponseWriter, r *http.Request) {
	auth.Logout(w, r)
	fmt.Println("Logout : " + structure.TplData.ProcessMessage)
	errTpl := tpl.ExecuteTemplate(w, "accueil.html", structure.TplData)
	if errTpl != nil {
		fmt.Println(errTpl)
	}

}

func SearchHandler(w http.ResponseWriter, r *http.Request) {
	input := r.FormValue("input")
	fmt.Println(input, "test input")
	http.Redirect(w, r, "/", http.StatusSeeOther)
}
