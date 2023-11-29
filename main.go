package main

import (
	o "OSINT/Back/server/handlers"
	"fmt"
	"net/http"
)

func main() {
	//all static files
	fs := http.FileServer(http.Dir("./Front/"))
	http.Handle("/Front/", http.StripPrefix("/Front/", fs))
	//all Handlers
	http.HandleFunc("/", o.WelcomeHandler)
	http.HandleFunc("/Recherche", o.SearchHandler)
	http.HandleFunc("/Enregistrement", o.EnregistrementHandler)
	http.HandleFunc("/Authentification", o.LoginHandler)
	// path := "path/to/script"
	// s.ScriptExec(path)
	//Start server
	fmt.Println("Serveur web écoutant sur le port 8080...")
	http.ListenAndServe(":8080", nil)

}
