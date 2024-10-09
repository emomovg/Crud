package db

import (
	"database/sql"
	_ "github.com/jackc/pgx/v4/stdlib"
	"log"
)

var DB *sql.DB

func Init() {
	var err error
	dsn := "postgres://app:pass@localhost:5432/db"
	DB, err = sql.Open("pgx", dsn)
	if err != nil {
		log.Println(err)
		return
	}
}
