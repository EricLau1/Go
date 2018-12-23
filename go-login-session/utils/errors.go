package utils

import (
	"net/http"
)

func InternalServerError(w http.ResponseWriter) {

	w.WriteHeader(http.StatusInternalServerError)
	w.Write([]byte("Erro interno do servidor..."))

}