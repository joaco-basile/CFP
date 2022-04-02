package api

import (
	"net/http"
)

type API struct{}

type calendario struct {
	id            int
	datos         string
	propietario   string
	colaboradores []string
}

func (a *API) getCalendario(w http.ResponseWriter, r *http.Request) {
}
