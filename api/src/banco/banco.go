package banco

import (
	"database/sql"

	"github.com/fabiokusaba/devbook/api/src/config"
	_ "github.com/go-sql-driver/mysql" // Driver
)

func Conectar() (*sql.DB, error) {
	db, erro := sql.Open("mysql", config.StringConexaoBanco)
	if erro != nil {
		return nil, erro
	}

	if erro = db.Ping(); erro != nil {
		db.Close()
		return nil, erro
	}

	return db, nil
}
