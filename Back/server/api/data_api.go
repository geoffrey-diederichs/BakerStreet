package api

import (
	"net/http"
)


func Extract_Api(research string) {
	logger.Info("tentativederecherche")
}

func Test_Api(w http.ResponseWriter, r *http.Request){
	Extract_Api("test")
	return
}