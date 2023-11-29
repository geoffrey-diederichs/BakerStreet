package authentification

import (
	data "OSINT/Back/server/data"
	structure "OSINT/Back/server/structure"
	"database/sql"
	"fmt"
	"log"
	"net/http"

	_ "github.com/mattn/go-sqlite3"

	"golang.org/x/crypto/bcrypt"
)

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
	_, ok := session.Values["pseudo"].(string)
	if ok {
		// Rediriger l'utilisateur vers la page d'acceuil s'il est connecté
		http.Redirect(w, r, "/acceuil", http.StatusSeeOther)
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
				structure.TplData.ProcessMessage = "Utilisateur inconnu"
				fmt.Println(structure.TplData.ProcessMessage)
				return // Ajout d'un return pour éviter de continuer l'exécution
			}
			// Autres erreurs
			log.Println("Erreur lors de la récupération du mot de passe :", err)
			structure.TplData.ProcessMessage = "Erreur interne"
			fmt.Println(structure.TplData.ProcessMessage)
			return
		}

		if err != nil {
			// retourne faux
			structure.TplData.ProcessMessage = "Mot de passe impossible à hasher"
			fmt.Println(structure.TplData.ProcessMessage)
			return
		}
		// si test est égal à hash

		// Vérifier que le mot de passe fourni correspond au hash stocké dans la base de données

		if ComparePassword(MdpHash, mdp) != nil {
			// Le mot de passe fourni ne correspond pas au hash stocké dans la base de données
			structure.TplData.ProcessMessage = "Mot de passe incorrect"
			fmt.Println(structure.TplData.ProcessMessage)
			return
		} else {
			structure.TplData.ProcessMessage = "Vous êtes connecté"
			fmt.Println(structure.TplData.ProcessMessage)
			return
		}

	} else {
		structure.TplData.ProcessMessage = "Entrez bien toute les informations"
		fmt.Println(structure.TplData.ProcessMessage)
		http.Redirect(w, r, "/", http.StatusSeeOther)
	}
}

func ComparePassword(hashPassword string, password string) error {
	pw := []byte(hashPassword)
	pass := []byte(password)
	return bcrypt.CompareHashAndPassword(pw, pass)
}
