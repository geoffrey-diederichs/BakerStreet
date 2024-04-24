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



type TemplateData struct {
	ProcessMessage string
	User           Utilisateurs
	History		   History
	Results        Results
}

type Results struct {
	facebook [] string
	tiktok [] string
	twitter [] string
	github [] string
}

type History struct {
	Yesterday [] string
	Last7Days [] string
	ThisMonth [] string
	Today [] string
}


