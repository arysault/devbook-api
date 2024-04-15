package repositories

import (
	"database/sql"
	"devbook-api/src/models"
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
		"insert into usuarios (nome, nick, email, senha) values(?, ?, ?, ?)",
	)

	if erro != nil {
		return 0, nil
	}

	defer statement.Close()

	result, erro := statement.Exec(usuario.Nome,
		usuario.Email, usuario.Nick, usuario.Senha)

	if erro != nil {
		return 0, nil
	}

	lastIdInserted, erro := result.LastInsertId()
	if erro != nil {
		return 0, erro
	}

	return uint64(lastIdInserted), nil
}
