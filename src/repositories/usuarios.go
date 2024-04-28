package repositories

import (
	"database/sql"
	"devbook-api/src/models"
	"fmt"
)

type usuarios struct {
	db *sql.DB
}

// NovoRepositorioDeUsuarios cria um repositorio de usuarios
func NovoRepositorioDeUsuarios(db *sql.DB) *usuarios {
	return &usuarios{db}
}

func (u usuarios) Criar(usuario models.Usuario) (uint64, error) {

	statement, erro := u.db.Prepare(
		"insert into usuarios(nome, nick, email, senha) values(?, ?, ?, ?)",
	)

	if erro != nil {
		return 0, erro
	}

	defer statement.Close()

	result, erro := statement.Exec(usuario.Nome,
		usuario.Nick, usuario.Email, usuario.Senha)

	if erro != nil {
		return 0, erro
	}

	lastIDInserted, erro := result.LastInsertId()
	if erro != nil {
		return 0, erro
	}

	return uint64(lastIDInserted), nil
}

// Buscar traz todos os usuarios atendendo o filtro de nome ou nick
func (u usuarios) Buscar(nomeOrNick string) ([]models.Usuario, error) {
	nomeOrNick = fmt.Sprintf("%%%s%%", nomeOrNick) //%nomeOrNick%

	rows, erro := u.db.Query(
		"select id, nome, nick, email, criadoEm from usuarios where nome LIKE ? or nick LIKE ?",
		nomeOrNick, nomeOrNick)
	if erro != nil {
		return nil, erro
	}
	defer rows.Close()

	var usuarios []models.Usuario

	for rows.Next() {
		var usuario models.Usuario
		if erro := rows.Scan(
			&usuario.ID,
			&usuario.Nome,
			&usuario.Nick,
			&usuario.Email,
			&usuario.CriadoEm,
		); erro != nil {
			return nil, erro
		}

		usuarios = append(usuarios, usuario)
	}

	return usuarios, nil
}

// FindByID traz um unico usuario pesquisado no banco de dados
func (u usuarios) FindByID(ID uint64) (models.Usuario, error) {
	rows, erro := u.db.Query(
		"select id, nome, nick, email, criadoEm from usuarios where id = ?",
		ID,
	)
	if erro != nil {
		return models.Usuario{}, erro
	}
	defer rows.Close()

	var user models.Usuario
	if rows.Next() {
		if erro = rows.Scan(
			&user.ID,
			&user.Nome,
			&user.Nick,
			&user.Email,
			&user.CriadoEm,
		); erro != nil {
			return models.Usuario{}, erro
		}

	}
	return user, nil
}
