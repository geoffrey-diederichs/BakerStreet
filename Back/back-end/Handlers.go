package OSINT

import (
	"fmt"
	"html/template"
	"net/http"
)

var tpl *template.Template

func init() {
	tpl = template.Must(template.ParseGlob("Front/*.html"))
}

func WelcomeHandler(w http.ResponseWriter, r *http.Request) {
	errTpl := tpl.ExecuteTemplate(w, "HomeTest.html", TplData)
	if errTpl != nil {
		fmt.Println(errTpl)
	}
}
