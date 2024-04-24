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

type Results struct {
	URL [] string
}

type TemplateData struct {
	ProcessMessage string
	User           Utilisateurs
	History		   History
	Results        Results
	// Result         Result
}

type History struct {
	Yesterday [] string
	Last7Days [] string
	ThisMonth [] string
	Today [] string
}


