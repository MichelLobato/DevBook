package modelos

import (
	"errors"
	"strings"
	"time"
)

// Usuario representa um usuario utilizando a rede social
type Usuario struct {
	ID       uint64    `json:"id,omitempty"`
	Nome     string    `json:"nome,omitempty"`
	Nick     string    `json:"nick,omitempty"`
	Email    string    `json:"email,omitempty"`
	Senha    string    `json:"senha,omitempty"`
	CriadoEm time.Time `json:"criadoem,omitempty"`
}

// Preparar vai chamar os metodos para validar e formatar o usuario recebido
func (usuario *Usuario) Preparar(etapa string) error {
	if erro := usuario.validar(etapa); erro != nil {
		return erro
	}

	usuario.formatar()
	return nil
}

func (usuario *Usuario) validar(etapa string) error {
	if usuario.Nome == "" {
		return errors.New("O nome e um campo obrigatorio nao pode estar em branco")
	}

	if usuario.Nick == "" {
		return errors.New("O nick e um campo obrigatorio nao pode estar em branco")
	}

	if usuario.Email == "" {
		return errors.New("O email e um campo obrigatorio nao pode estar em branco")
	}

	if etapa == "cadastro" && usuario.Senha == "" {
		return errors.New("O senha e um campo obrigatorio nao pode estar em branco")
	}

	return nil
}

func (usuario *Usuario) formatar() {
	usuario.Nome = strings.TrimSpace(usuario.Nome)
	usuario.Nick = strings.TrimSpace(usuario.Nick)
	usuario.Email = strings.TrimSpace(usuario.Email)
}
