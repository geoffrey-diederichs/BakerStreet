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
	tpl = template.Must(template.ParseGlob("Front/HomeTest.html"))
}

func WelcomeHandler(w http.ResponseWriter, r *http.Request) {
	structure.TplData.ProcessMessage = "Welcome to the OSINT project !"
	errTpl := tpl.ExecuteTemplate(w, "HomeTest.html", structure.TplData)
	if errTpl != nil {
		fmt.Println(errTpl)
	}
}

func EnregistrementHandler(w http.ResponseWriter, r *http.Request) {
	auth.Enregistrement(w, r)
	fmt.Println("Enregistrement : " + structure.TplData.ProcessMessage)
	errTpl := tpl.ExecuteTemplate(w, "HomeTest.html", structure.TplData)
	if errTpl != nil {
		fmt.Println(errTpl)
	}
}

func LoginHandler(w http.ResponseWriter, r *http.Request) {
	auth.Login(w, r)
	fmt.Println("Login : " + structure.TplData.ProcessMessage)
	errTpl := tpl.ExecuteTemplate(w, "HomeTest.html", structure.TplData)
	if errTpl != nil {
		fmt.Println(errTpl)
	}

}

func LogoutHandler(w http.ResponseWriter, r *http.Request) {
	auth.Logout(w, r)
	fmt.Println("Logout : " + structure.TplData.ProcessMessage)
	errTpl := tpl.ExecuteTemplate(w, "HomeTest.html", structure.TplData)
	if errTpl != nil {
		fmt.Println(errTpl)
	}

}

func SearchHandler(w http.ResponseWriter, r *http.Request) {
	input := r.FormValue("input")
	fmt.Println(input, "test input")
	http.Redirect(w, r, "/", http.StatusSeeOther)
}
