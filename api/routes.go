package api

import (
	"net/http"

	"github.com/gorilla/mux"
)

func (a *API) RegisterRoutes(r *mux.Router) {
	r.HandleFunc("/calendario", a.postCalendario).Methods(http.MethodPost)
	r.HandleFunc("/calendario", a.getCalendario).Methods(http.MethodGet)
	r.HandleFunc("/calendario", a.patchCalendario).Methods(http.MethodPatch)
	//r.HandleFunc("/calendario", a.deleteCalendario).Methods(http.MethodDelete)

}
