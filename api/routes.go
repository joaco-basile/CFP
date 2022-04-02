package api

import (
	"github.com/gorilla/mux"
)

func (a *API) RegisterRoutes(r *mux.Router) {
	r.HandleFunc("/calendario", a.getCalendario).Methods("GET")
}
