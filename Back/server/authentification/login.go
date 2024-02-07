package authentification

import (
	data "OSINT/Back/server/data"
	logs "OSINT/Back/server/logs"
	structure "OSINT/Back/server/structure"
	"database/sql"
	"net/http"

	_ "github.com/mattn/go-sqlite3"
	"go.uber.org/zap"

	"golang.org/x/crypto/bcrypt"
)

var logger = logs.GetLog(logs.GetLogConfig())

// check virifie les identifiant de l'uilisateur
func Login(w http.ResponseWriter, r *http.Request) {

	// recuperation de de la de la session utilisateur
	session, err := data.Store.Get(r, "data")

	// gestion de l'erreur
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	// Vérifier si l'utilisateur est connecté
	_, ok := session.Values["username"].(string)
	if ok {

		// Rediriger l'utilisateur vers la page d'acceuil s'il est connecté
		structure.TplData.ProcessMessage = "Vous êtes déja connecté en tant que " + session.Values["username"].(string) + " !"
		http.Redirect(w, r, "/", http.StatusSeeOther)
		return
	}

	if r.Method == "POST" {
		username := r.FormValue("username")
		mdp := r.FormValue("password")
		// Récupérer le hash du mot de passe enregistré dans la base de données pour cet utilisateur
		row := data.Bd.QueryRow("SELECT mdp FROM Utilisateurs WHERE username = ?", username)

		var MdpHash string // ou []byte si le hash est stocké sous forme binaire
		err := row.Scan(&MdpHash)

		if err != nil {
			if err == sql.ErrNoRows {
				// L'utilisateur n'existe pas dans la base de données
				logger.Info("Utilisateur inconnu")
				return // Ajout d'un return pour éviter de continuer l'exécution
			}
			// Autres erreurs
			logger.Error("Erreur lors de la récupération du mot de passe :", zap.Error(err))
			return
		}

		if err != nil {
			// retourne faux
			logger.Error("Mot de passe impossible à hasher", zap.Error(err))
			return
		}
		// si test est égal à hash

		// Vérifier que le mot de passe fourni correspond au hash stocké dans la base de données

		if ComparePassword(MdpHash, mdp) != nil {
			// Le mot de passe fourni ne correspond pas au hash stocké dans la base de données
			structure.TplData.ProcessMessage = "Mot de passe incorrect"
			return
		} else {
			session.Values["username"] = username
			session.Save(r, w)
			structure.TplData.ProcessMessage = "Vous êtes maintenant connecté en tant que " + username + " !"
			logger.Info("User" + username + "logged in successfully")
			http.Redirect(w, r, "/", http.StatusSeeOther)
		}

	} else {
		structure.TplData.ProcessMessage = "Entrez bien toute les informations"

	}
}

func ComparePassword(hashPassword string, password string) error {
	pw := []byte(hashPassword)
	pass := []byte(password)
	return bcrypt.CompareHashAndPassword(pw, pass)
}
