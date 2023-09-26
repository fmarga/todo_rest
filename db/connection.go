package db

import (
	"database/sql"
	"fmt"
	"todo_rest/configs"

	_ "github.com/lib/pq"
)

func OpenConnection() (*sql.DB, error) {
	//somente abre conexao com banco
	conf := configs.GetDB()

	sc := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=disable", conf.Host, conf.Port, conf.User, conf.Pass, conf.Database)

	conn, err := sql.Open("postgres", sc)
	if err != nil {
		panic(err) // nao e boa pratica utilizar em prod
	}

	err = conn.Ping() // testa a conexao

	return conn, err
}
