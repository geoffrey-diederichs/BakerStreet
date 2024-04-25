package api

import (
	 structure "OSINT/Back/server/structure"
	"bufio"
	"encoding/json"
	"fmt"
	"net/http"
	"os"
	// "path/filepath"

	"go.uber.org/zap"
)

type Results struct {
	Facebook []string `json:"facebook"`
	TikTok   []string `json:"tiktok"`
	Twitter  []string `json:"twitter"`
	GitHub   []string `json:"github"`
}

	var Facebook []string
    var Github []string
    var Twitter []string
    var Tiktok []string



func Extract_Api(research string) (Results) {
	

    // path := "/Osint/app/api/api.txt"

    var results Results
    // absPath, err := filepath.Abs(path)
    // if err != nil {
    //     logger.Error("Erreur lors de l'obtention du chemin absolu", zap.Error(err))
    //     return results
    // }
    // logger.Info("Chemin absolu du fichier:", zap.String("path", absPath))


    file, err := os.Open(abs_path)
    if err != nil {
        logger.Error("Erreur lors de l'ouverture du fichier", zap.Error(err))
        return results
    }
    defer file.Close()

    
    scanner := bufio.NewScanner(file)


    
    if scanner.Scan() {
        line := scanner.Text()
        
    
        err := json.Unmarshal([]byte(line), &results)
        if err != nil {
            logger.Error("Erreur lors du décodage de la ligne JSON", zap.Error(err))
            return results
        }
    } else {
        logger.Error("Impossible de lire la première ligne du fichier")
        return results
    }

    if err := scanner.Err(); err != nil {
        logger.Error("Erreur lors de la lecture du fichier", zap.Error(err))
        return results
    }
    fmt.Println("Facebook:") 

	for _, link := range results.Facebook {
		fmt.Println(link)
        Facebook = append(Facebook, link)
	}
    structure.TplData.Results.Facebook = Facebook

	fmt.Println("TikTok:")
	for _, link := range results.TikTok {
        Tiktok = append(Tiktok, link)
		fmt.Println(link)
	}
    structure.TplData.Results.Tiktok = Tiktok

	fmt.Println("Twitter:")
	for _, link := range results.Twitter {
		fmt.Println(link)
        Twitter = append(Twitter, link)
	}
    structure.TplData.Results.Twitter = Twitter

	fmt.Println("GitHub:")
	for _, link := range results.GitHub {
     Github = append(Github, link)
		fmt.Println(link)
	}
    structure.TplData.Results.Github = Github
 // Affiche à nouveau la structure TplData après l'appel à Extract_Api
    fmt.Printf("TplData après l'appel à Extract_Api : %+v\n", structure.TplData.Results)

    return results
}


func Test_Api(w http.ResponseWriter, r *http.Request){
    Extract_Api("test")
    return
}