package db

import (
	"context"
	"github.com/jackc/pgx/v4/pgxpool"
	"log"
	"time"
)

var Pool *pgxpool.Pool

func Init() {
	var err error
	dsn := "postgres://app:pass@localhost:5432/db"
	connectCtx, _ := context.WithTimeout(context.Background(), 5*time.Second)
	Pool, err = pgxpool.Connect(connectCtx, dsn)
	if err != nil {
		log.Println(err)
		return
	}
}
