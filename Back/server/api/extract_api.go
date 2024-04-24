package api

import (
	logs "OSINT/Back/server/logs"
	"bufio"
	"os"
	"path/filepath"

	"go.uber.org/zap"
)

var logger = logs.GetLog(logs.GetLogConfig())

func Extract_Api(research string) {
	// L'URL pour récupérer le premier post
	path := "/Osint/app/api/api.txt"
	abs_path, err := filepath.Abs(path)
	if err != nil {
        logger.Info("Error getting absolute path:", zap.Error(err))
        return
    }
    logger.Info("Absolute path:", zap.String(abs_path,abs_path))

	file, err := os.Open(path)
	if err != nil {
		logger.Info("Error opening file:", zap.Error(err))
		return
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	if scanner.Scan() {
		logger.Info("First line:", zap.String(scanner.Text(),scanner.Text()))
	}

	if err := scanner.Err(); err != nil {
		logger.Info("Error reading file:", zap.Error(err))
	}
}
