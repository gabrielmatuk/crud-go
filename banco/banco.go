package banco

import (
	"crud/credenciais"
	"database/sql"

	_ "github.com/go-sql-driver/mysql" //Driver de conexão
)

func Conectar() (*sql.DB, error) {
	credenciais := credenciais.Credenciais("DB_STRING")
	stringConexao := credenciais

	db, erro := sql.Open("mysql", stringConexao)

	if erro != nil {
		return nil, erro
	}

	if erro = db.Ping(); erro != nil {
		return nil, erro
	}

	return db, nil
}
