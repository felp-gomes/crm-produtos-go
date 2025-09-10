package db

import (
	"database/sql"

	_ "github.com/lib/pq"
)


func ConectaComBancoDeDados() *sql.DB {
	conexao := "user=postgres dbname=crm-produto-go password=timecon@passWord123 host=localhost port=5432 sslmode=disable"
	db, err := sql.Open("postgres", conexao)
	if err != nil {
		panic(err.Error())
	}
	return db
}