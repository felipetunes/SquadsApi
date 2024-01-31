package db

import (
	"database/sql"

	_ "github.com/go-sql-driver/mysql"
)

// Passa as informações sobre para a conexão com o banco
const database = "mysql"
const user = "admin"
const password = "Tunes1313#"
const server = "teammate.cr2mw0ioqij3.us-east-1.rds.amazonaws.com"
const door = "3306"
const databaseName = "Teammate"

// Função que conecta ao banco de dados
func ConnectDB() (*sql.DB, error) {

	// Abre a conexão com o banco de dados
	db, err := sql.Open(database, user+":"+password+"@tcp("+server+":"+door+")/"+databaseName)
	if err != nil {
		return nil, err
	}

	// Testa se a conexão está ativa
	err = db.Ping()
	if err != nil {
		return nil, err
	}

	// Define o fuso horário para São Paulo
	_, err = db.Exec("SET time_zone = 'America/Sao_Paulo'")
	if err != nil {
		return nil, err
	}

	// Retorna a conexão como resultado
	return db, nil
}
