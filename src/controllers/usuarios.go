package controllers

import "net/http"

//CriarUsiarios cria insere usuaios no Banco
func CriarUsuarios(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Criando usuários"))
}

//BuscarUsuarios busca uma lista de usuarios no banco de dados
func BuscarUsuarios(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Buscando usuários"))
}

//BuscarUsuario busca um usuário por ID no banco
func BuscarUsuario(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Buscando usuário"))
}

//AtualizarUsuario atualiza o usuario no banco
func AtualizarUsuario(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Atualizando usuários"))
}

//DeletarUsuario deleta usuário no banco
func DeletarUsuario(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Deletando usuários"))
}
