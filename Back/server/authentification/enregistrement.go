package authentification

import (
	data "OSINT/Back/server/data"
	structure "OSINT/Back/server/structure"
	"net/http"
	"net/mail"
	"strconv"

	"go.uber.org/zap"
)

func isEmailValid(e string) bool {
	_, err := mail.ParseAddress(e)
	return err == nil
}

func Enregistrement(w http.ResponseWriter, r *http.Request) {
	logger.Info("You are tryng to register")
	session, err := data.Store.Get(r, "data")
	if err != nil {
		logger.Error("Failed to get the session : ", zap.Error(err))
	}
	_, ok := session.Values["username"].(string)
	logger.Info("session")
	if ok {
		logger.Info("Vous êtes déja connecté en tant que " + session.Values["username"].(string) + " !")
		structure.TplData.ProcessMessage = "Vous êtes déja connecté en tant que " + session.Values["username"].(string) + " !"
		http.Redirect(w, r, "/", http.StatusSeeOther)
		return
	}
	
	if r.Method == "POST" {
		logger.Info("You are tryng to register POST")
		nom := r.FormValue("nom")
		prenom := r.FormValue("prenom")
		password := r.FormValue("password")
		email := r.FormValue("email")
		confirme_password := r.FormValue("confirme_password")
		username := r.FormValue("username")
		ageStr := r.FormValue("age")
		icon := r.FormValue("icon")
		logger.Info("User registration", zap.String("username", username), zap.String("email", email), zap.String("nom", nom), zap.String("prenom", prenom), zap.String("age", ageStr), zap.String("icon", icon), zap.String("password", password), zap.String("confirme_password", confirme_password))
		if username == "" || email == "" || password == "" || confirme_password == "" || nom == "" || prenom == "" || ageStr == "" {
			logger.Info("Entrez bien toute les informations")
			structure.TplData.ProcessMessage = "Entrez bien toute les informations"
			return
		}
		if password != confirme_password {
			structure.TplData.ProcessMessage = "Mot de passe non identique"
			logger.Info("Password not identical")
			return
		}
		// isPasswordValid, msgPasswordVerification := validatePassword(password)
		// if !isPasswordValid {
		// 	structure.TplData.ProcessMessage = msgPasswordVerification
		// 	return
		// }
		// if len(username) < 4 {
		// 	structure.TplData.ProcessMessage = "Nom d'utilisateur trop court"
		// 	return
		// }
		// if len(nom) < 2 {
		// 	structure.TplData.ProcessMessage = "Nom trop court"
		// 	return
		// }
		// if len(prenom) < 2 {
		// 	structure.TplData.ProcessMessage = "Prenom trop court"
		// 	return
		// }

		age, err := strconv.Atoi(ageStr)
		if err != nil {
			structure.TplData.ProcessMessage = "Merci d'entrer un age valide"
			logger.Error("Age is not an integer", zap.String("age", ageStr), zap.Error(err))
			return
		}
		
		// if !isEmailValid(username + "<" + email +">") {
		// 	structure.TplData.ProcessMessage = "Email invalide"
		// 	return
		// }

		// token,err := GenerateToken()
		// if err != nil {
		// 	structure.TplData.ProcessMessage = "Erreur le mail de vérification n'a pas pu être envoyé, veuillez réessayer l'inscription"
		// 	return
		// }
		
		// if sendVerificationEmail(email, token) != nil {	
		// 	structure.TplData.ProcessMessage = "Erreur le mail de vérification n'a pas pu être envoyé, veuillez réessayer l'inscription"
		// 	return
		// }

		mdpHashed, err := HashPassword(password)
		if err != nil {
			logger.Error("password can't be hashed", zap.Error(err))
		}
		_, errAddUser := data.Bd.Exec("INSERT INTO Utilisateurs (username, mdp,nom,prenom,email,age,icon) VALUES (?, ?, ?, ?, ?,?,?)", username, mdpHashed, nom, prenom, email, age, icon)

		if errAddUser != nil {
			switch errAddUser.Error() {
			case "UNIQUE constraint failed: Utilisateurs.email":
				structure.TplData.ProcessMessage = "email déja utilisé" + errAddUser.Error()
				logger.Debug("email déja utilisé", zap.Error(errAddUser))
			case "UNIQUE constraint failed: Utilisateurs.username":
				structure.TplData.ProcessMessage = "nom utilisateur déja utilisé" + errAddUser.Error()
				logger.Debug("nom déja utilisé", zap.Error(errAddUser))
			default:
				structure.TplData.ProcessMessage = "Sql error : " + errAddUser.Error() + "\n"
				logger.Debug("Sql error : ", zap.Error(errAddUser))
			}
			logger.Info("User registration failed")
			return
		} else {
			structure.TplData.ProcessMessage = "Tu es maintenant inscrit en tant que " + username + " !"
			session.Values["username"] = username
			session.Save(r, w)
			logger.Info("User " + username + " registered successfully")
			http.Redirect(w, r, "/", http.StatusSeeOther)
		}

	} else {
		logger.Info("Bienvenue sur la page d'enregistrement")
		return
	}

}
