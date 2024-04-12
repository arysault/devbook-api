package main

import (
	"devbook-api/src/config"
	"devbook-api/src/router"
	"fmt"
	"log"
	"net/http"
)

func main() {
	config.Carregar()
	fmt.Println(config.StringConexaoBanco)
	fmt.Println(config.Porta)
	fmt.Println("Rodando a API")

	r := router.Gerar()

	log.Fatal(http.ListenAndServe(
		fmt.Sprintf(":%d", config.Porta), r))
}
