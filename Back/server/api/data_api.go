package api

import (
	 structure "OSINT/Back/server/structure"
	"bufio"
	"encoding/json"
    	"os"
	"go.uber.org/zap"
)


func Extract_Api() {

    var results structure.Results

    file, err := os.Open(abs_path)
    if err != nil {
        logger.Error("Erreur lors de l'ouverture du fichier", zap.Error(err))
        return 
    }
    defer file.Close()

    
    scanner := bufio.NewScanner(file)
    
    if scanner.Scan() {
        line := scanner.Text()
        err := json.Unmarshal([]byte(line), &results)
        if err != nil {
            logger.Error("Erreur lors du décodage de la ligne JSON", zap.Error(err))
            return 
        }
    } else {
        logger.Error("Impossible de lire la première ligne du fichier")
        return 
    }

    if err := scanner.Err(); err != nil {
        logger.Error("Erreur lors de la lecture du fichier", zap.Error(err))
        return 
    }
    structure.TplData.Results = results
 // Affiche à nouveau la structure TplData après l'appel à Extract_Api
    // fmt.Printf("TplData après l'appel à Extract_Api : %+v\n", structure.TplData.Results)

}
