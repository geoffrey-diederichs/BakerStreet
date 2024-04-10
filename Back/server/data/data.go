package data

import (
	logs "OSINT/Back/server/logs"
	"database/sql"

	"github.com/gorilla/sessions"
	_ "github.com/mattn/go-sqlite3"
	"go.uber.org/zap"
)

var logger = logs.GetLog(logs.GetLogConfig())

// Var pour définir la base de donée
var Bd, err = OuvrirBaseDonnee("./data/db.sqlite")

// Store est var de la cle pour les cookies
var Store = sessions.NewCookieStore([]byte("Motdepassesupersecurisealamortquitu"))

// OuvrirBaseDonnee est une fonction pour ouvrir la connexion à la base de donnée et creer les tables si elles n'existent pas
func OuvrirBaseDonnee(chemin string) (*sql.DB, error) {
	bd, err := sql.Open("sqlite3", chemin)
	if err != nil {
		logger.Error("Failed to open db : ", zap.Error(err))
	}
	logger.Info("connexion a la base de donnée réussite")
	_, err = bd.Exec("SELECT * FROM Utilisateurs")
	if err != nil {
		// Si la table n'existe pas, la créer
		_, err := bd.Exec(`CREATE TABLE Utilisateurs (
    id INTEGER PRIMARY KEY AUTOINCREMENT,
    username TEXT UNIQUE NOT NULL,
    mdp TEXT NOT NULL,
	nom TEXT ,
	prenom TEXT ,
	email TEXT UNIQUE NOT NULL,
	age INT,
	icon TEXT
);`)
		if err != nil {
			logger.Error("Failed to create db", zap.Error(err))
			return bd, err
		}
		logger.Info("Table Utilisateurs créée avec succès.")
	} else {
		logger.Info("La table Utilisateurs existe déjà.")
	}

	_, err = bd.Exec(`CREATE TABLE History (
        id INT AUTO_INCREMENT PRIMARY KEY,
        username TEXT NOT NULL,
        research VARCHAR(255) NOT NULL,
        timestamp DATETIME NOT NULL DEFAULT CURRENT_TIMESTAMP
        );`)
        if err != nil {
            logger.Error("Failed to create history table", zap.Error(err))
            return bd, err
        }

	return bd, err
}

// ObtenirInfoUtilisateur est une fonction pour avoir les informations de l'utilisateur demandé
func ObtenirInfoUtilisateur(NomUtilisateur string) (int, string, string, string, int, string, error) {
	// Initialisation de demande avec une requette SQL pour recuperer les informations de l'utilisateur
	demande := Bd.QueryRow("SELECT id, prenom, nom, mail, age, icon FROM Utilisateurs WHERE pseudo = ?", NomUtilisateur)
	// Initialisation des variables id, prenom, nom, mail, age et icon
	var id int
	var prenom string
	var nom string
	var email string
	var age int
	var icon string
	// On rentre les informations de demande dans les variables
	err := demande.Scan(&id, &prenom, &nom, &email, &age, &icon)
	// S'il y a une erreur renvoie des variables contenant rien plus l'erreur
	if err != nil {
		return 0, "", "", "", 0, "", err
	}
	// Sinon envoyer les variables
	return id, prenom, nom, email, age, icon, nil
}

func ObtenirInfoUtilisateurID(idUtilisateur int) (int, string, string, string, int, string, error) {
	// Initialisation de demande avec une requette SQL pour recuperer les informations de l'utilisateur
	demande := Bd.QueryRow("SELECT id, prenom, nom, email, age, icon FROM Utilisateurs WHERE id = ?", idUtilisateur)
	// Initialisation des variables id, prenom, nom, mail, age et icon
	var id int
	var prenom string
	var nom string
	var email string
	var age int
	var icon string
	// On rentre les informations de demande dans les variables
	err := demande.Scan(&id, &prenom, &nom, &email, &age, &icon)
	// S'il y a une erreur renvoie des variables contenant rien plus l'erreur
	if err != nil {
		return 0, "", "", "", 0, "", err
	}
	// Sinon envoyer les variables
	return id, prenom, nom, email, age, icon, nil
}

func ObtenirInfoPoste(ID string) (string, string, string, string, string, int, int, error) {
	// Initialisation de demande avec une requette SQL pour recuperer les informations de l'utilisateur
	demande := Bd.QueryRow("SELECT theme, titre, description, cree_le, cree_par, likes, dislikes FROM Postes WHERE id = ?", ID)
	// Initialisation des variables id, prenom, nom, mail, age et icon
	var theme, titre, description, cree_le, cree_par string
	var likes, dislikes int
	// On rentre les informations de demande dans les variables
	err := demande.Scan(&theme, &titre, &description, &cree_le, &cree_par, likes, dislikes)
	// S'il y a une erreur renvoie des variables contenant rien plus l'erreur
	if err != nil {
		return "", "", "", "", "", 0, 0, err
	}
	// Sinon envoyer les variables
	return theme, titre, description, cree_le, cree_par, likes, dislikes, nil
}
