package api

import (
	"database/sql"
	"net/http"

	_ "github.com/go-sql-driver/mysql"
)

type API struct{}

type calendario struct {
	id            int
	datos         string
	propietario   string
	colaboradores []string
}

func (a *API) getCalendario(w http.ResponseWriter, r *http.Request) {
	db, err := sql.Open("mysql", "roo:etec@tcp(localhost:3306)/calendario")

	if err != nil {
		panic(err.Error())
	}

}
