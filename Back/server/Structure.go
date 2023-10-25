package OSINT

var TplData TemplateData

type User struct {
	Id    string
	Name  string
	Email string
}

// type Result struct {
// 	Pseudo string
// 	Email string
// Location string
// 	Phone string

// }

type TemplateData struct {
	ProcessMessage string
	User           User
	// Result         Result
}
