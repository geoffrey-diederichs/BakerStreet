package api

import (
	structure "OSINT/Back/server/structure"
	"bufio"
	"encoding/json"
	"os"
	"regexp"
	"time"

	"go.uber.org/zap"
)

func Extract_Results() string {
	time.Sleep(15 * time.Second)
	var lastLine string
	file, err := os.Open(abs_path)
	if err != nil {
		logger.Error("Error opening file:", zap.Error(err))
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
		if match != nil {
			if len(match) == 3 {
				logger.Info("Extracting Results : operation successful")
				return match[2]
			} else {
				logger.Error("Extracting Results : incorrect result line in api.txt file")
				return ""
			}
		} else {
			logger.Error("Extracting Results : No input found")
			return ""
		}

	}

	if err := scanner.Err(); err != nil {
		logger.Error("Error reading file:", zap.Error(err))
	}
	return ""
}

func Convert_Json(line string) {

	var results structure.Results

	err := json.Unmarshal([]byte(line), &results)

	if err != nil {
		logger.Error("Erreur lors du décodage de la ligne JSON", zap.Error(err))
		return
	}
	logger.Info("All results added to page after converting JSON")
	structure.TplData.Results = results

}
