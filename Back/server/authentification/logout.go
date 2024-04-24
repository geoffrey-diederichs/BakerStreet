package authentification

import (
	data "OSINT/Back/server/data"
	structure "OSINT/Back/server/structure"
	"net/http"
)

func Logout(w http.ResponseWriter, r *http.Request) {
    // Retrieve the user session
    session, err := data.Store.Get(r, "data")
    if err != nil {
        http.Error(w, err.Error(), http.StatusInternalServerError)
        return
    }

    // Invalidate the session
    session.Values = map[interface{}]interface{}{} // Clear all data in the session
    session.Options.MaxAge = -1
	structure.TplData.User.Username = ""
	structure.TplData.User.Nom = ""
	structure.TplData.User.Prenom = ""
	structure.TplData.User.Email = ""
	structure.TplData.User.Age = 0
	structure.TplData.User.Icon = ""

    // Save the changes to the session
    if err := session.Save(r, w); err != nil {
        http.Error(w, err.Error(), http.StatusInternalServerError)
        return
    }

    // Log and inform the user
    logger.Info("Vous êtes maintenant déconnecté")
    structure.TplData.ProcessMessage = "Vous êtes maintenant déconnecté"
    http.Redirect(w, r, "/", http.StatusSeeOther)
}

