package OSINT

import (
	"fmt"
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
		nom := r.FormValue("nom")
		prenom := r.FormValue("prenom")
		password := r.FormValue("password")
		email := r.FormValue("email")
		confirme_password := r.FormValue("confirme_password")
		username := r.FormValue("username")
		if username == "" || email == "" || password == "" || confirme_password == "" || nom == "" || prenom == "" {
			TplData.ProcessMessage = "Entrez bien toute les informations"
			fmt.Println(TplData.ProcessMessage)
			http.Redirect(w, r, "/", http.StatusSeeOther)
			return
		} else if password != confirme_password {
			TplData.ProcessMessage = "Mot de passe non identique"
			fmt.Println(TplData.ProcessMessage)
			http.Redirect(w, r, "/", http.StatusSeeOther)
			return
		}
		mdpHashed, err := HashPassword(password)
		if err != nil {
			print(err)
		}

		_, errAddUser := Bd.Exec("INSERT INTO Utilisateurs (username, mdp,nom,prenom,email) VALUES (?, ?, ?, ?, ?)", username, mdpHashed, nom, prenom, email)

		if errAddUser != nil {
			switch errAddUser.Error() {
			case "UNIQUE constraint failed: Utilisateurs.email":
				TplData.ProcessMessage = "email déja utilisé"
				fmt.Println(TplData.ProcessMessage)
			case "UNIQUE constraint failed: Utilisateurs.username":
				TplData.ProcessMessage = "nom utilisateur déja utilisé"
				fmt.Println(TplData.ProcessMessage)
			default:
				TplData.ProcessMessage = "Sql error : " + errAddUser.Error() + "\n"
				fmt.Println(TplData.ProcessMessage)
			}

			http.Redirect(w, r, "/", http.StatusSeeOther)
			return
		} else {
			TplData.ProcessMessage = "You have been registered. Please log in"
			fmt.Println(TplData.ProcessMessage)
			http.Redirect(w, r, "/", http.StatusSeeOther)
		}

	} else {
		TplData.ProcessMessage = "Entrez bien toute les informations"
		fmt.Println(TplData.ProcessMessage)
		http.Redirect(w, r, "/", http.StatusSeeOther)
	}

}
