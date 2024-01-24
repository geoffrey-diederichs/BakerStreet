package authentification

import (
	data "OSINT/Back/server/data"
	structure "OSINT/Back/server/structure"
	"net/http"

	"go.uber.org/zap"
)

func Enregistrement(w http.ResponseWriter, r *http.Request) {
	session, err := data.Store.Get(r, "data")
	if err != nil {
		logger.Error("Failed to get the session : ", zap.Error(err))
		return
	}
	_, ok := session.Values["username"].(string)
	if ok {
		structure.TplData.ProcessMessage = "Vous êtes déja connecté en tant que " + session.Values["username"].(string) + " !"
		http.Redirect(w, r, "/", http.StatusSeeOther)
		return
	}
	if r.Method == "POST" {

		nom := r.FormValue("nom")
		prenom := r.FormValue("prenom")
		password := r.FormValue("password")
		email := r.FormValue("email")
		confirme_password := r.FormValue("confirme_password")
		username := r.FormValue("username")
		age := r.FormValue("age")
		icon := r.FormValue("icon")
		if username == "" || email == "" || password == "" || confirme_password == "" || nom == "" || prenom == "" || age == "" {
			logger.Info("Entrez bien toute les informations")
			structure.TplData.ProcessMessage = "Entrez bien toute les informations"
			return
		} else if password != confirme_password {
			structure.TplData.ProcessMessage = "Mot de passe non identique"
			return
		}
		mdpHashed, err := HashPassword(password)
		if err != nil {
			logger.Error("password can't be hashed", zap.Error(err))
		}

		_, errAddUser := data.Bd.Exec("INSERT INTO Utilisateurs (username, mdp,nom,prenom,email,age,icon) VALUES (?, ?, ?, ?, ?,?,?)", username, mdpHashed, nom, prenom, email, age, icon)

		if errAddUser != nil {
			switch errAddUser.Error() {
			case "UNIQUE constraint failed: Utilisateurs.email":
				structure.TplData.ProcessMessage = "email déja utilisé" + errAddUser.Error()
			case "UNIQUE constraint failed: Utilisateurs.username":
				structure.TplData.ProcessMessage = "nom utilisateur déja utilisé" + errAddUser.Error()
			default:
				structure.TplData.ProcessMessage = "Sql error : " + errAddUser.Error() + "\n"
				logger.Debug("Sql error : ", zap.Error(errAddUser))
			}

			return
		} else {
			structure.TplData.ProcessMessage = "Tu es mainentenant inscrit en tant que " + username + " !"
			session.Values["username"] = username
			session.Save(r, w)
			logger.Info("User " + username + " registered successfully")
			http.Redirect(w, r, "/", http.StatusSeeOther)
		}

	} else {
		structure.TplData.ProcessMessage = "Entrez bien toute les informations"
		return
	}

}
