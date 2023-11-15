package OSINT

import (
	"net/http"

	_ "github.com/mattn/go-sqlite3"
)

func Enregistrement(w http.ResponseWriter, r *http.Request) {
	println("test")
	if r.Method == "POST" {

		// if r.FormValue("submit") == "" {
		// 	TplData.ProcessMessage = "Entrez bien toute les informations"
		// 	http.Redirect(w, r, "/", http.StatusSeeOther)
		// } else {
		println("test")
		password := r.FormValue("password")
		email := r.FormValue("email")
		username := r.FormValue("username")
		if username == "" || email == "" || password == "" {
			TplData.ProcessMessage = "Entrez bien toute les informations"
		}
		mdpHashed, err := HashPassword(password)
		if err != nil {
			print(err)
		}

		_, errAddUser := Bd.Exec("INSERT INTO Utilisateurs (pseudo, mdp) VALUES (?, ?)", username, mdpHashed)

		// fmt.Println(email)
		// fmt.Println(username)
		// fmt.Println(passwordHashed)

		taken, _ := PseudoCheck(username)
		if taken {
			messageerror := "Utilisateur deja utilisé !"
			TplData.ProcessMessage = messageerror
			http.Redirect(w, r, "/", http.StatusSeeOther)
		}

		if errAddUser != nil {
			switch errAddUser.Error() {
			case "UNIQUE constraint failed: users.email":
				TplData.ProcessMessage = "This email is already in use"
			case "UNIQUE constraint failed: users.name":
				TplData.ProcessMessage = "This username is already in use"
			case "CHECK constraint failed: name <> ''":
				TplData.ProcessMessage = "Please enter an username"
			case "CHECK constraint failed: email <> ''":
				TplData.ProcessMessage = "Please enter an email"
			}

			http.Redirect(w, r, "/", http.StatusSeeOther)
			return
		} else {
			TplData.ProcessMessage = "You have been registered. Please log in"
			http.Redirect(w, r, "/", http.StatusSeeOther)
		}

	} else {
		TplData.ProcessMessage = "Entrez bien toute les informations"
		http.Redirect(w, r, "/", http.StatusSeeOther)
	}

}
