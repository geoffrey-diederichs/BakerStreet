package api

import (
	 structure "OSINT/Back/server/structure"
	"bufio"
	"encoding/json"
    "regexp"
    "os"
    "time"
	"go.uber.org/zap"
)

func Extract_Results() string {
    time.Sleep(9 * time.Second)
	var lastLine string
	file, err := os.Open(abs_path)
	if err != nil {
		logger.Info("Error opening file:", zap.Error(err))
        return ""
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		lastLine = scanner.Text()
	}

	if lastLine != "" {
		re := regexp.MustCompile(`^(.*?);(.+)$`)
		match := re.FindStringSubmatch(lastLine)
        if(match != nil){
            if len(match) == 3 {
                return match[2]
            } else {
                logger.Info("incorrect result line in api.txt file")
                return ""
            }
        }else{
            logger.Info("No data found after")
            return ""
        }

	}

	if err := scanner.Err(); err != nil {
		logger.Info("Error reading file:", zap.Error(err))
	}
	return ""
}

func Convert_Json (line string) {

    var results structure.Results
    
    err := json.Unmarshal([]byte(line), &results)

    if err != nil {
        logger.Error("Erreur lors du décodage de la ligne JSON", zap.Error(err))
        return 
    }

    structure.TplData.Results = results
    // Affiche à nouveau la structure TplData après l'appel à Extract_Api
    // fmt.Printf("TplData après l'appel à Extract_Api : %+v\n", structure.TplData.Results)

}
