package OSINT

import (
	"database/sql"
	"fmt"
	"log"
	"net/http"

	_ "github.com/mattn/go-sqlite3"

	"golang.org/x/crypto/bcrypt"
)

// check virifie les identifiant de l'uilisateur
func Login(w http.ResponseWriter, r *http.Request) {

	if r.Method == "POST" {
		username := r.FormValue("username")
		mdp := r.FormValue("password")
		// Récupérer le hash du mot de passe enregistré dans la base de données pour cet utilisateur
		row := Bd.QueryRow("SELECT mdp FROM Utilisateurs WHERE username = ?", username)

		var MdpHash string // ou []byte si le hash est stocké sous forme binaire
		err := row.Scan(&MdpHash)

		if err != nil {
			if err == sql.ErrNoRows {
				// L'utilisateur n'existe pas dans la base de données
				TplData.ProcessMessage = "Utilisateur inconnu"
				fmt.Println(TplData.ProcessMessage)
				http.Redirect(w, r, "/", http.StatusSeeOther)
				return // Ajout d'un return pour éviter de continuer l'exécution
			}
			// Autres erreurs
			log.Println("Erreur lors de la récupération du mot de passe :", err)
			TplData.ProcessMessage = "Erreur interne"
			fmt.Println(TplData.ProcessMessage)
			http.Redirect(w, r, "/", http.StatusSeeOther)
			return
		}

		if err != nil {
			// retourne faux
			TplData.ProcessMessage = "Mot de passe impossible à hasher"
			fmt.Println(TplData.ProcessMessage)
			http.Redirect(w, r, "/", http.StatusSeeOther)
		}
		// si test est égal à hash

		// Vérifier que le mot de passe fourni correspond au hash stocké dans la base de données

		if ComparePassword(MdpHash, mdp) != nil {
			// Le mot de passe fourni ne correspond pas au hash stocké dans la base de données
			TplData.ProcessMessage = "Mot de passe incorrect"
			fmt.Println(TplData.ProcessMessage)
			http.Redirect(w, r, "/", http.StatusSeeOther)
		} else {
			TplData.ProcessMessage = "Vous êtes connecté"
			fmt.Println(TplData.ProcessMessage)
			http.Redirect(w, r, "/", http.StatusSeeOther)
		}

	} else {
		TplData.ProcessMessage = "Entrez bien toute les informations"
		fmt.Println(TplData.ProcessMessage)
		http.Redirect(w, r, "/", http.StatusSeeOther)
	}
}

func ComparePassword(hashPassword string, password string) error {
	pw := []byte(hashPassword)
	pass := []byte(password)
	return bcrypt.CompareHashAndPassword(pw, pass)
}
