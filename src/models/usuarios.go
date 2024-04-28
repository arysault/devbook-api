package models

import (
	"errors"
	"strings"
	"time"
)

// Usuario representa um user
type Usuario struct {
	ID       uint64    `json:"id,omitempty"`
	Nome     string    `json:"nome,omitempty"`
	Nick     string    `json:"nick,omitempty"`
	Email    string    `json:"email,omitempty"`
	Senha    string    `json:"senha,omitempty"`
	CriadoEm time.Time `json:"CriadoEm,omitempty"`
}

// Preparar vai chamar os metodos validar em formatar de usuário
func (usuario *Usuario) Preparar() error {
	if erro := usuario.validar(); erro != nil {
		return erro
	}

	usuario.formatar()
	return nil
}

func (usuario *Usuario) validar() error {

	if usuario.Nome == "" {
		return errors.New("o usuário não pode ficar em branco")
	}

	if usuario.Nick == "" {
		return errors.New("o Nick não pode ficar em branco")
	}

	if usuario.Email == "" {
		return errors.New("o Email não pode ficar em branco")
	}

	if usuario.Senha == "" {
		return errors.New("a senha não pode ficar em branco")
	}

	return nil
}

func (usuario *Usuario) formatar() {
	usuario.Nome = strings.TrimSpace(usuario.Nome)
	usuario.Nick = strings.TrimSpace(usuario.Nick)
	usuario.Email = strings.TrimSpace(usuario.Email)
}
