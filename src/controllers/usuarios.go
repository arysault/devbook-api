package controllers

import (
	"devbook-api/src/banco"
	"devbook-api/src/models"
	"devbook-api/src/repositories"
	"devbook-api/src/responses"
	"encoding/json"
	"io/ioutil"
	"net/http"
	"strconv"
	"strings"

	"github.com/gorilla/mux"
)

// CriarUsiarios cria insere usuaios no Banco
func CriarUsuarios(w http.ResponseWriter, r *http.Request) {
	corpoRequest, erro := ioutil.ReadAll(r.Body)
	if erro != nil {
		responses.Erro(w, http.StatusUnprocessableEntity, erro)
		return
	}

	var usuario models.Usuario
	if erro = json.Unmarshal(corpoRequest, &usuario); erro != nil {
		responses.Erro(w, http.StatusBadRequest, erro)
		return
	}

	if erro = usuario.Preparar(); erro != nil {
		responses.Erro(w, http.StatusBadRequest, erro)
		return
	}

	db, erro := banco.Conectar()
	if erro != nil {
		responses.Erro(w, http.StatusInternalServerError, erro)
		return
	}

	defer db.Close()

	repositories := repositories.NovoRepositorioDeUsuarios(db)
	usuario.ID, erro = repositories.Criar(usuario)
	if erro != nil {
		responses.Erro(w, http.StatusInternalServerError, erro)
		return
	}

	responses.JSON(w, http.StatusCreated, usuario)

}

// BuscarUsuarios busca uma lista de usuarios no banco de dados
func BuscarUsuarios(w http.ResponseWriter, r *http.Request) {
	nomeOrNick := strings.ToLower(r.URL.Query().Get("usuario"))
	db, erro := banco.Conectar()
	if erro != nil {
		responses.Erro(w, http.StatusInternalServerError, erro)
		return
	}

	defer db.Close()

	repositories := repositories.NovoRepositorioDeUsuarios(db)
	usuarios, erro := repositories.Buscar(nomeOrNick)
	if erro != nil {
		responses.Erro(w, http.StatusInternalServerError, erro)
		return
	}
	responses.JSON(w, http.StatusOK, usuarios)
}

// BuscarUsuario busca um usuário por ID no banco
func BuscarUsuario(w http.ResponseWriter, r *http.Request) {
	paramaters := mux.Vars(r)

	userID, erro := strconv.ParseUint(paramaters["usuarioId"], 10, 64)
	if erro != nil {
		responses.Erro(w, http.StatusBadRequest, erro)
		return
	}

	db, erro := banco.Conectar()
	if erro != nil {
		responses.Erro(w, http.StatusInternalServerError, erro)
		return
	}
	defer db.Close()

	repositories := repositories.NovoRepositorioDeUsuarios(db)
	user, erro := repositories.FindByID(userID)
	if erro != nil {
		responses.Erro(w, http.StatusInternalServerError, erro)
		return
	}
	responses.JSON(w, http.StatusOK, user)
}

// AtualizarUsuario atualiza o usuario no banco
func AtualizarUsuario(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Atualizando usuários"))
}

// DeletarUsuario deleta usuário no banco
func DeletarUsuario(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Deletando usuários"))
}
