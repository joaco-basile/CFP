package api

import (
	"database/sql"
	"fmt"
	"log"
	"net/http"

	_ "github.com/go-sql-driver/mysql"
)

type API struct{}

type calendario struct {
	id                                        int
	nombre, propietario, datos, colaboradores string
}

type calendarios []calendario

func openDb() (db *sql.DB) {
	db, err := sql.Open("mysql", "root:etec@tcp(localhost:3306)/calendarios")
	if err != nil {
		panic("fallo la conexion con la base")
	}

	return db
}

func (a *API) postCalendario(w http.ResponseWriter, r *http.Request) {
	c := calendario{
		id:          1,
		nombre:      "disponibilidad de camiones",
		propietario: "fran",
		datos:       "'dias,camion1,camion2,camion3\nlunes,horario1,horario2,horario3'",
	}

	db := openDb()

	sentenciaPreparada, err1 := db.Prepare("INSERT INTO calendario (nombre, propietario, datos, colaboradores) VALUES(?, ?, ?, ?)")
	if err1 != nil {
		log.Fatal(err1)
	}
	defer sentenciaPreparada.Close()

	_, err2 := sentenciaPreparada.Exec(c.nombre, c.propietario, c.datos, c.colaboradores)

	if err2 != nil {
		panic("fallo la execucion de la centencia")
	}
}

func (a *API) getCalendario(w http.ResponseWriter, r *http.Request) {
	var calendarios calendarios
	var c calendario
	db := openDb()
	rows, err := db.Query("SELECT * FROM calendario")
	if err != nil {
		log.Fatal(err)
	}
	defer rows.Close()
	for rows.Next() {
		err := rows.Scan(&c.id, &c.nombre, &c.propietario, &c.datos, &c.colaboradores)
		if err != nil {
			log.Fatal(err)
		}
		calendarios = append(calendarios, c)
	}
	fmt.Println(calendarios)

}

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
