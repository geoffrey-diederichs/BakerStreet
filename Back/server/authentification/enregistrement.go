package authentification

import (
	data "OSINT/Back/server/data"
	structure "OSINT/Back/server/structure"
	"fmt"
	"net/http"

	_ "github.com/mattn/go-sqlite3"
)

func Enregistrement(w http.ResponseWriter, r *http.Request) {
	session, err := data.Store.Get(r, "data")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	_, ok := session.Values["pseudo"].(string)
	if ok {
		http.Redirect(w, r, "/accueil", http.StatusSeeOther)
		return
	}
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
			structure.TplData.ProcessMessage = "Entrez bien toute les informations"
			fmt.Println(structure.TplData.ProcessMessage)
			return
		} else if password != confirme_password {
			structure.TplData.ProcessMessage = "Mot de passe non identique"
			fmt.Println(structure.TplData.ProcessMessage)
			return
		}
		mdpHashed, err := HashPassword(password)
		if err != nil {
			print(err)
		}

		_, errAddUser := data.Bd.Exec("INSERT INTO Utilisateurs (username, mdp,nom,prenom,email) VALUES (?, ?, ?, ?, ?)", username, mdpHashed, nom, prenom, email)

		if errAddUser != nil {
			switch errAddUser.Error() {
			case "UNIQUE constraint failed: Utilisateurs.email":
				structure.TplData.ProcessMessage = "email déja utilisé"
				fmt.Println(structure.TplData.ProcessMessage)
			case "UNIQUE constraint failed: Utilisateurs.username":
				structure.TplData.ProcessMessage = "nom utilisateur déja utilisé"
				fmt.Println(structure.TplData.ProcessMessage)
			default:
				structure.TplData.ProcessMessage = "Sql error : " + errAddUser.Error() + "\n"
				fmt.Println(structure.TplData.ProcessMessage)
			}

			return
		} else {
			structure.TplData.ProcessMessage = "You have been registered. Please log in"
			fmt.Println(structure.TplData.ProcessMessage)
			return
		}

	} else {
		structure.TplData.ProcessMessage = "Entrez bien toute les informations"
		fmt.Println(structure.TplData.ProcessMessage)
		return
	}

}
