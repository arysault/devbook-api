package controllers

import (
	"devbook-api/src/banco"
	"devbook-api/src/models"
	"devbook-api/src/repositories"
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"
)

// CriarUsiarios cria insere usuaios no Banco
func CriarUsuarios(w http.ResponseWriter, r *http.Request) {
	corpoRequest, erro := ioutil.ReadAll(r.Body)
	if erro != nil {
		log.Fatal(erro)
	}

	var usuario models.Usuario
	if erro = json.Unmarshal(corpoRequest, &usuario); erro != nil {
		log.Fatal(erro)
	}

	db, erro := banco.Conectar()
	if erro != nil {
		log.Fatal(erro)
	}

	repositories := repositories.NovoRepositorioDeUsuarios(db)
	repositories.Criar(usuario)
}

// BuscarUsuarios busca uma lista de usuarios no banco de dados
func BuscarUsuarios(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Buscando usuários"))
}

// BuscarUsuario busca um usuário por ID no banco
func BuscarUsuario(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Buscando usuário"))
}

// AtualizarUsuario atualiza o usuario no banco
func AtualizarUsuario(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Atualizando usuários"))
}

// DeletarUsuario deleta usuário no banco
func DeletarUsuario(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Deletando usuários"))
}
