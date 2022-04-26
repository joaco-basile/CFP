package api

import (
	"database/sql"
	"log"
	"net/http"
	"strconv"

	_ "github.com/go-sql-driver/mysql"
	"github.com/labstack/echo"
)

type API struct{}
type Calendario struct {
	ID            int    `json:"idCaledario"`
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

func (a *API) getCalendario(ec echo.Context) error {

	var c Calendario
	var cs Calendarios

	db := openDb()
	defer db.Close()

	id := ec.QueryParams().Get("id")

	idInt, err := strconv.Atoi(id)

	if err != nil {
		panic(err)
	}

	rows, err := db.Query("SELECT * FROM calendario WHERE idCalendario = ?", idInt)
	if err != nil {
		log.Fatal(err)
	}
	defer rows.Close()

	for rows.Next() {
		err := rows.Scan(&c.ID, &c.Nombre, &c.Propietario, &c.Datos, &c.Colaboradores)
		if err != nil {
			log.Fatal(err)
		}
		cs = append(cs, c)
	}

	return ec.JSON(http.StatusOK, cs)
}
func (a *API) getCalendarios(ec echo.Context) error {

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
		err := rows.Scan(&c.ID, &c.Nombre, &c.Propietario, &c.Datos, &c.Colaboradores)
		if err != nil {
			log.Fatal(err)
		}
		cs = append(cs, c)
	}

	return ec.JSON(http.StatusOK, cs)
}

func (a *API) postCalendario(ec echo.Context) error {

	c := Calendario{
		Nombre:        ec.QueryParams().Get("nombre"),
		Propietario:   ec.QueryParams().Get("propietario"),
		Datos:         ec.QueryParams().Get("datos"),
		Colaboradores: ec.QueryParams().Get("colaboradores"),
	}

	db := openDb()
	defer db.Close()

	sentenciaPreparada, err1 := db.Prepare("INSERT INTO calendario (nombre, propietario, datos, colaboradores) VALUES(?, ?, ?, ?)")

	if err1 != nil {
		log.Fatal(err1)
	}
	defer sentenciaPreparada.Close()

	_, err2 := sentenciaPreparada.Exec(c.Nombre, c.Propietario, c.Datos, c.Colaboradores)

	if err2 != nil {
		panic("fallo la execucion de la centencia")
	}

	return ec.JSON(http.StatusAccepted, map[string]string{"mensaje": "El usuario se registro con exito"})
}
func (a *API) patchCalendario(ec echo.Context) error {

	c := Calendario{
		Nombre:        ec.QueryParams().Get("nombre"),
		Propietario:   ec.QueryParams().Get("propietario"),
		Datos:         ec.QueryParams().Get("datos"),
		Colaboradores: ec.QueryParams().Get("colaboradores"),
	}

	var err error
	c.ID, err = strconv.Atoi(ec.QueryParams().Get("id"))
	if err != nil {
		log.Fatal(err)
	}

	db := openDb()
	defer db.Close()

	sentenciaPreparada, err1 := db.Prepare("UPDATE calendario SET nombre = ?, propietario = ?, datos = ?, colaboradores = ? WHERE idCalendario = ?")
	if err1 != nil {
		log.Fatal(err1)
	}
	defer sentenciaPreparada.Close()

	result, err2 := sentenciaPreparada.Exec(c.Nombre, c.Propietario, c.Datos, c.Colaboradores, c.ID)

	filasAfectadas, _ := result.RowsAffected()

	if filasAfectadas == 0 {
		return ec.JSON(http.StatusBadRequest, "no se encontro el valor que se quiere modificar")
	}

	if err2 != nil {
		log.Fatal("fallo la execucion de la centencia")
	}
	return ec.JSON(http.StatusAccepted, map[string]string{"mensaje": "se cambiaron los datos"})
}
func (a *API) deleteCalendario(ec echo.Context) error {
	nombre := ec.QueryParams().Get("nombre")

	db := openDb()
	defer db.Close()

	sentenciaPreparada, err := db.Prepare("DELETE FROM calendario WHERE nombre = ?")
	if err != nil {
		log.Fatal(err)
	}
	defer sentenciaPreparada.Close()

	result, err := sentenciaPreparada.Exec(nombre)

	filasAfectadas, _ := result.RowsAffected()

	if filasAfectadas == 0 {
		return ec.JSON(http.StatusBadRequest, "no se encontro el valor que se quiere eliminar")
	}

	if err != nil {
		log.Fatal(err)
	}

	return ec.JSON(http.StatusAccepted, map[string]string{"mensaje": "se elimino el o los items de la tabla calendario"})
}
