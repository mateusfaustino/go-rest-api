package db

import (
	"database/sql"
	"fmt"
	"log"

	_ "github.com/go-sql-driver/mysql"
)

const (
	host     = "localhost"
	port     = 3307
	user     = "root"
	password = ""
	dbname   = "go_rest_api"
)

func ConnectDb() (*sql.DB, error) {
	// Define a string de conexão
	dsn := fmt.Sprintf("%s:%s@tcp(%s)/%s", user, password, host, dbname)

	// Abre a conexão com o banco de dados
	db, err := sql.Open("mysql", dsn)
	if err != nil {
		return nil, err
	}

	// Verifica se a conexão está funcionando corretamente
	err = db.Ping()
	if err != nil {
		return nil, err
	}

	log.Println("Conexão com o banco de dados estabelecida com sucesso")
	return db, nil
}
