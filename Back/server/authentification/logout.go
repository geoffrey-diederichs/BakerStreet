package authentification

import (
	data "OSINT/Back/server/data"
	structure "OSINT/Back/server/structure"
	"fmt"
	"net/http"
)

// Deconnexion gere la deconnexion de l'utilisateur
func Logout(w http.ResponseWriter, r *http.Request) {

	// recuperation de de la de la session utilisateur
	session, err := data.Store.Get(r, "data")

	// gestion de l'erreur
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	// Supprime la session
	session.Options.MaxAge = -1
	err = session.Save(r, w)

	// gestion de l'erreur
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	// redirection de l'utilisateur vers l'acceuil
	structure.TplData.ProcessMessage = "Vous êtes maintenant déconnecté"
	fmt.Println(structure.TplData.ProcessMessage)
	return
}
