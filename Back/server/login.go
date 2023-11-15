package OSINT

import (
	"database/sql"
	"log"
	"net/http"

	_ "github.com/mattn/go-sqlite3"

	"golang.org/x/crypto/bcrypt"
)

// check virifie les identifiant de l'uilisateur
func Check(w http.ResponseWriter, r *http.Request) {

	if r.Method == "POST" {
		// pseudo := r.FormValue("username")
		pseudo := "tim"
		mdp := r.FormValue("password")
		// Récupérer le hash du mot de passe enregistré dans la base de données pour cet utilisateur
		row := Bd.QueryRow("SELECT mdp FROM Utilisateurs WHERE pseudo = ?", pseudo)

		var MdpHash string // ou []byte si le hash est stocké sous forme binaire
		err := row.Scan(&MdpHash)

		if err != nil {
			if err == sql.ErrNoRows {
				// L'utilisateur n'existe pas dans la base de données
				TplData.ProcessMessage = "Utilisateur inconnu"
				http.Redirect(w, r, "/", http.StatusSeeOther)
				return // Ajout d'un return pour éviter de continuer l'exécution
			}
			// Autres erreurs
			log.Println("Erreur lors de la récupération du mot de passe :", err)
			TplData.ProcessMessage = "Erreur interne"
			http.Redirect(w, r, "/", http.StatusSeeOther)
			return
		}

		if err != nil {
			// retourne faux
			TplData.ProcessMessage = "Mot de passe impossible à hasher"
			http.Redirect(w, r, "/", http.StatusSeeOther)
		}
		// si test est égal à hash

		// Vérifier que le mot de passe fourni correspond au hash stocké dans la base de données

		if ComparePassword(MdpHash, mdp) != nil {
			println("false")
			// Le mot de passe fourni ne correspond pas au hash stocké dans la base de données
			TplData.ProcessMessage = "Mot de passe incorrect"
			http.Redirect(w, r, "/", http.StatusSeeOther)
		} else {
			println("true")
			TplData.ProcessMessage = "Vous êtes connecté"
			http.Redirect(w, r, "/", http.StatusSeeOther)
		}

	} else {
		TplData.ProcessMessage = "Entrez bien toute les informations"
		http.Redirect(w, r, "/", http.StatusSeeOther)
	}
}

func ComparePassword(hashPassword string, password string) error {
	pw := []byte(hashPassword)
	pass := []byte(password)
	return bcrypt.CompareHashAndPassword(pw, pass)
}
