package app

import (
	"mycrudapp/internal/db"
	http2 "mycrudapp/internal/handlers/http"
	"mycrudapp/internal/repo"
	"mycrudapp/internal/routes"
	"net/http"
)

func Run() {
	pg, err := db.Init()
	if err != nil {
		return
	}

	defer pg.Pool.Close()

	r := routes.NewRouter()
	customerRepository := repo.NewCustomerRepository(pg.Pool)
	customerHandler := http2.NewCustomerHandler(customerRepository)
	r.RegisterRoutes(customerHandler)
	errHttp := http.ListenAndServe(":8080", r.MuxRouter)
	if errHttp != nil {
		return
	}

}
