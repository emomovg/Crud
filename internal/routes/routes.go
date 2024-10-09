package routes

import (
	"github.com/gorilla/mux"
	customerHandlers "mycrudapp/internal/handlers/http"
	"net/http"
)

func NewRouter() *mux.Router {
	router := mux.NewRouter()

	router.HandleFunc("/customers", customerHandlers.GetAll).Methods(http.MethodGet)
	router.HandleFunc("/customers/{id:[0-9]+}", customerHandlers.GetById).Methods(http.MethodGet)
	router.HandleFunc("/customers", customerHandlers.Create).Methods(http.MethodPost)
	router.HandleFunc("/customers/{id:[0-9]+}", customerHandlers.Update).Methods(http.MethodPut)
	router.HandleFunc("/customers/{id:[0-9]+}", customerHandlers.Delete).Methods(http.MethodDelete)
	router.HandleFunc("/customers/{id:[0-9]+}/activate", customerHandlers.Activate).Methods(http.MethodPut)
	router.HandleFunc("/customers/{id:[0-9]+}/deactivate", customerHandlers.Deactivate).Methods(http.MethodPut)
	router.HandleFunc("/customers/getAllActivated", customerHandlers.GetAllActivated).Methods(http.MethodGet)

	return router
}
