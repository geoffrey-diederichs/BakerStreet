package main

import (
	logs "OSINT/Back/server/logs"
	"encoding/json"
	"fmt"
	"net/http"
	"os"
	"go.uber.org/zap/zapcore"
)

var logger = logs.GetLog(logs.GetLogConfig())

func main() {
	// L'URL pour récupérer le premier post
	url := "https://jsonplaceholder.typicode.com/posts/1"

	// Effectuer la requête HTTP GET
	resp, err := http.Get(url)
	if err != nil {
		logger.Info("Erreur lors de la requête HTTP : %s\n", zapcore.Error(err))
		os.Exit(1)
	}
	defer resp.Body.Close()

	// Décoder la réponse JSON dans une instance de Post
	// var post Post
	if err := json.NewDecoder(resp.Body).Decode(&post); err != nil {
		fmt.Printf("Erreur lors du décodage du JSON : %s\n", err)
		os.Exit(1)
	}

	// Afficher le titre du post
	fmt.Println("Titre du Post :", post.Title)
}
