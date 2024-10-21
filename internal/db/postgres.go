package db

import (
	"context"
	"github.com/jackc/pgx/v4/pgxpool"
	"log"
	"time"
)

type Postgres struct {
	Pool *pgxpool.Pool
}

func Init() (*Postgres, error) {
	var err error
	dsn := "postgres://app:pass@localhost:5432/db"
	connectCtx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	pg, err := pgxpool.Connect(connectCtx, dsn)
	if err != nil {
		log.Println(err)
		return nil, err
	}

	return &Postgres{Pool: pg}, nil
}
