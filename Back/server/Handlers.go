package OSINT

import (
	"fmt"
	"html/template"
	"net/http"
)

var tpl *template.Template

func init() {
	tpl = template.Must(template.ParseGlob("Front/HomeTest.html"))
}

func WelcomeHandler(w http.ResponseWriter, r *http.Request) {
	TplData.ProcessMessage = "Welcome to the OSINT project !"
	errTpl := tpl.ExecuteTemplate(w, "HomeTest.html", TplData)
	if errTpl != nil {
		fmt.Println(errTpl)
	}
}

func SearchHandler(w http.ResponseWriter, r *http.Request) {
	input := r.FormValue("input")
	fmt.Println(input, "test input")
	// s.ScriptExec(input)
	http.Redirect(w, r, "/", http.StatusSeeOther)
}
