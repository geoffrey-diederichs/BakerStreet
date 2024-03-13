package structure

import "time"

var TplData TemplateData

type Utilisateurs struct {
	ID       int
	Username string
	Mdp      string
	Prenom   string
	Nom      string
	Email    string
	Age      int
	Icon     string
}

type Histories struct {
	ID        int
	Research  string
	Timestamp time.Time
}

// type Result struct {
// 	Pseudo string
// 	Email string
// Location string
// 	Phone string
// }

type TemplateData struct {
	ProcessMessage string
	User           Utilisateurs
	History        Histories
	// Result         Result
}
