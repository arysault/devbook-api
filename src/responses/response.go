package responses

import (
	"encoding/json"
	"log"
	"net/http"
)

// JSON retorna uma resposta em JSON a requisição
func JSON(w http.ResponseWriter, statusCode int, data interface{}) {
	w.WriteHeader(statusCode)

	if erro := json.NewEncoder(w).Encode(data); erro != nil {
		log.Fatal(erro)
	}
}

// Erro retorna um erro em formato JSON
func Erro(w http.ResponseWriter, statusCode int, erro error) {
	JSON(w, statusCode, struct {
		Error string `json:"error"`
	}{Error: erro.Error()})
}
