package repositorios

import (
	"database/sql"

	"github.com/fabiokusaba/devbook/api/src/modelos"
)

type Usuarios struct {
	db *sql.DB
}

func NovoRepositorioDeUsuarios(db *sql.DB) *Usuarios {
	return &Usuarios{
		db: db,
	}
}

func (u Usuarios) Criar(usuario modelos.Usuario) (uint64, error) {
	return 0, nil
}
