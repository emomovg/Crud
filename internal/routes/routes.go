package routes

import (
	"github.com/gorilla/mux"
	customerHandlers "mycrudapp/internal/handlers/http"
	"net/http"
)

type Router struct {
	MuxRouter *mux.Router
}

func NewRouter() *Router {
	return &Router{
		MuxRouter: mux.NewRouter(),
	}
}

func (r *Router) RegisterRoutes(handler *customerHandlers.CustomerHandler) {

	r.MuxRouter.HandleFunc("/customers", handler.GetAll).Methods(http.MethodGet)
	r.MuxRouter.HandleFunc("/customers/{id:[0-9]+}", handler.GetById).Methods(http.MethodGet)
	r.MuxRouter.HandleFunc("/customers", handler.Create).Methods(http.MethodPost)
	r.MuxRouter.HandleFunc("/customers/{id:[0-9]+}", handler.Update).Methods(http.MethodPut)
	r.MuxRouter.HandleFunc("/customers/{id:[0-9]+}", handler.Delete).Methods(http.MethodDelete)
	r.MuxRouter.HandleFunc("/customers/{id:[0-9]+}/activate", handler.Activate).Methods(http.MethodPut)
	r.MuxRouter.HandleFunc("/customers/{id:[0-9]+}/deactivate", handler.Deactivate).Methods(http.MethodPut)
	r.MuxRouter.HandleFunc("/customers/getAllActivated", handler.GetAllActivated).Methods(http.MethodGet)
}
