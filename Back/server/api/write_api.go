package api

import (
	logs "OSINT/Back/server/logs"
	"bufio"
	"fmt"
	"os"
	"path/filepath"
	"regexp"

	"go.uber.org/zap"
)

var logger = logs.GetLog(logs.GetLogConfig())

var abs_path = findPath()

func findPath() string {
	path := "../BakerStreet/Osint/app/api/api.txt"
	abs_path, err := filepath.Abs(path)
	if err != nil {
		logger.Info("Error getting absolute path:", zap.Error(err))
		return ""
	}
	logger.Info("Absolute path:", zap.String(abs_path, abs_path))
	return abs_path
}

func isSearching() bool {
	var lastLine string
	file, err := os.Open(abs_path)
	if err != nil {
		logger.Info("Error opening file:", zap.Error(err))
		return true
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
				logger.Info("No incomplete research was found, starting new research now...")
				return false
            } else {
                logger.Info("cant write to file : incorrect result line in api.txt file")
				return true
			}
        }else{
            logger.Info("An incomplete research was found, try later")
        }

	}

	if err := scanner.Err(); err != nil {
		logger.Info("Error reading file:", zap.Error(err))
	}
	return false
}

func Write_Api(research string) {
	if !isSearching() {
		file, err := os.OpenFile(abs_path, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
		if err != nil {
			logger.Info("Error opening file:", zap.Error(err))
			return
		}
		defer file.Close()
		research += ";"
		_, err = fmt.Fprintln(file, research)
		if err != nil {
			logger.Info("Error writing to file:", zap.Error(err))
		}
		logger.Info("",zap.String("Input user written to file",research))
	} else {
		logger.Error("Cant write new input to file : old input already being searched")
	}
}
