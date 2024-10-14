package app

import (
	"mycrudapp/internal/db"
	"mycrudapp/internal/routes"
	"net/http"
)

func Run() {
	db.Init()

	defer db.Pool.Close()

	r := routes.NewRouter()

	err := http.ListenAndServe(":8080", r)
	if err != nil {
		return
	}

}
