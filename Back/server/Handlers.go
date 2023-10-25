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

func RegistrationHandler(w http.ResponseWriter, r *http.Request) {
	if r.FormValue("submit") == "" {
		http.Redirect(w, r, "/", http.StatusSeeOther)
	} else {
		// processMessage := ""
		password := r.FormValue("password")
		email := r.FormValue("email")
		username := r.FormValue("username")
		if username == "" || email == "" || password == "" {
			http.Error(w, "All fields are required", http.StatusBadRequest)
			return
		}
		passwordHashed := HashPassword(password)
		fmt.Println(email)
		fmt.Println(username)
		fmt.Println(passwordHashed)
		// errAddUSer := addUser(username, email, passwordHashed)
		// if errAddUSer != nil {
		// 	switch errAddUSer.Error() {
		// 	case "UNIQUE constraint failed: users.email":
		// 		processMessage = "This email is already in use"
		// 	case "UNIQUE constraint failed: users.name":
		// 		processMessage = "This username is already in use"
		// 	case "CHECK constraint failed: name <> ''":
		// 		processMessage = "Please enter an username"
		// 	case "CHECK constraint failed: email <> ''":
		// 		processMessage = "Please enter an email"
		// 	}
		// 	fmt.Println(errAddUSer)
		// 	TplData.ProcessMessage = processMessage
		// 	http.Redirect(w, r, "/", http.StatusSeeOther)
		// } else {
		// 	processMessage = "You have been registered. Please log in"
		// 	TplData.ProcessMessage = processMessage
		// 	http.Redirect(w, r, "/", http.StatusSeeOther)
		// }
	}
}
