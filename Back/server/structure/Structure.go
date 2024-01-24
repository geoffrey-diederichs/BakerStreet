package structure

var TplData TemplateData

type Utilisateurs struct {
	ID     int
	Username string
	Mdp    string
	Prenom string
	Nom    string
	Email   string
	Age    int
	Icon   string
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
	// Result         Result
}
