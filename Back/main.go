package main

import (
	o "OSINT/back-end"
	"fmt"
	"net/http"
)

func main() {
	fs := http.FileServer(http.Dir("./Front/"))
	http.Handle("/Front/", http.StripPrefix("/Front/", fs))

	http.HandleFunc("/", o.WelcomeHandler)

	fmt.Println("Serveur web écoutant sur le port 8080...")
	http.ListenAndServe(":8080", nil)

}
