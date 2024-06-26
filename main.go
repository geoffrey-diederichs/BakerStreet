package main

import (
	o "OSINT/Back/server/handlers"
	"fmt"
	"net/http"
)

func main() {
	//all static files
	fs := http.FileServer(http.Dir("Front"))
	http.Handle("/Front/", http.StripPrefix("/Front/", fs))
	//all Handlers
	http.HandleFunc("/", o.WelcomeHandler)
	http.HandleFunc("/Recherche", o.SearchHandler)
	http.HandleFunc("/Resultats", o.ApiHandler)
	http.HandleFunc("/Enregistrement", o.EnregistrementHandler)
	// http.HandleFunc("/PasswordModify", o.PasswordModifyHandler)
	http.HandleFunc("/Authentification", o.LoginHandler)
	http.HandleFunc("/Account", o.AccountHandler)
	http.HandleFunc("/Deconnexion", o.LogoutHandler)
	// path := "path/to/script"
	// s.ScriptExec(path)
	//Start server
	fmt.Println("Serveur web écoutant sur le port 8080...")
	http.ListenAndServe(":8080", nil)

}
