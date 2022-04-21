package api

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	_ "github.com/go-sql-driver/mysql"
)

type API struct{}

type Calendario struct {
	Id            int    `json:"id"`
	Nombre        string `json:"nombre"`
	Propietario   string `json:"propietario"`
	Datos         string `json:"datos"`
	Colaboradores string `json:"colaboradores"`
}

type Calendarios []Calendario

func openDb() (db *sql.DB) {
	db, err := sql.Open("mysql", "root:etec@tcp(localhost:3306)/calendarios")
	if err != nil {
		panic("fallo la conexion con la base")
	}

	return db
}

func (a *API) postCalendario(w http.ResponseWriter, r *http.Request) {

	c := Calendario{
		Nombre:        "joaco",
		Propietario:   "joaco",
		Datos:         "nada",
		Colaboradores: "",
	}

	db := openDb()

	sentenciaPreparada, err1 := db.Prepare("INSERT INTO calendario (nombre, propietario, datos, colaboradores) VALUES(?, ?, ?, ?)")
	if err1 != nil {
		log.Fatal(err1)
	}
	defer sentenciaPreparada.Close()

	_, err2 := sentenciaPreparada.Exec(c.Nombre, c.Propietario, c.Datos, c.Colaboradores)

	if err2 != nil {
		panic("fallo la execucion de la centencia")
	}
}

func (a *API) getCalendario(w http.ResponseWriter, r *http.Request) {

	var c Calendario
	var cs Calendarios

	db := openDb()
	defer db.Close()

	rows, err := db.Query("SELECT * FROM calendario")
	if err != nil {
		log.Fatal(err)
	}
	defer rows.Close()
	for rows.Next() {
		err := rows.Scan(&c.Id, &c.Nombre, &c.Propietario, &c.Datos, &c.Colaboradores)
		if err != nil {
			log.Fatal(err)
		}
		cs = append(cs, c)
	}

	json.NewEncoder(w).Encode(cs)
}

// for que arme un array de strings
func (a *API) patchCalendario(w http.ResponseWriter, r *http.Request) {

	datosPeticion := r.URL.Query()["id"]

	fmt.Println(datosPeticion)
	/*
		c := calendario{
			id:          datosPeticion["id"][1],
			nombre:      datosPeticion.nombre,
			propietario: datosPeticion.propietario,
			datos:       datosPeticion.datos,
		}

		db := openDb()
		defer db.Close()

		sentenciaPreparada, err1 := db.Prepare("UPDATE calendario SET nombre = ?, propietario = ?, datos = ?, colaboradores = ? WHERE idCalendario = ?")
		if err1 != nil {
			log.Fatal(err1)
		}
		defer sentenciaPreparada.Close()

		_, err2 := sentenciaPreparada.Exec(c.nombre, c.propietario, c.datos, c.colaboradores)

		if err2 != nil {
			panic("fallo la execucion de la centencia")
		}
	*/
}
