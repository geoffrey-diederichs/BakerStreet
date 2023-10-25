package main

import (
	o "OSINT/Back/server"
	"fmt"
	"net/http"
)

func main() {
	//all static files
	fs := http.FileServer(http.Dir("./Front/"))
	http.Handle("/Front/", http.StripPrefix("/Front/", fs))
	//all Handlers
	http.HandleFunc("/", o.WelcomeHandler)
	//Start server
	fmt.Println("Serveur web écoutant sur le port 8080...")
	http.ListenAndServe(":8080", nil)

}
