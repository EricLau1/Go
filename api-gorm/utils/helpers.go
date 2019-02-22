package utils

import (
	"log"
	"net/http"
	"encoding/json"
)

func ToError(w http.ResponseWriter, err error, statusCode int) {
	ToJson(w, struct{
		Message string `json:"message"` 
	}{
		Message: err.Error(),
	}, statusCode)
}

func ToJson(w http.ResponseWriter, data interface{}, statusCode int) {
	w.Header().Set("Content-type", "application/json")
	w.WriteHeader(statusCode)
	err := json.NewEncoder(w).Encode(data)
	CheckError(err)
}

func CheckError(err error) {
	if err != nil {
		log.Fatal(err)
	}
}
