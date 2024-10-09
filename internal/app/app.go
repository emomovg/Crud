package app

import (
	"mycrudapp/internal/db"
	"mycrudapp/internal/routes"
	"net/http"
)

func Run() {
	db.Init()
	defer func() {
		if err := db.DB.Close(); err != nil {
			return
		}
	}()

	r := routes.NewRouter()

	err := http.ListenAndServe(":8080", r)
	if err != nil {
		return
	}

}
