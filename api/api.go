package api

import (
	"database/sql"
	"log"
	"net/http"
	"strconv"

	_ "github.com/go-sql-driver/mysql"
	"github.com/labstack/echo"
)

//Inicializando los tipos
type (
	API        struct{}
	Calendario struct {
		ID            int    `json:"idCaledario"`
		Nombre        string `json:"nombre"`
		Propietario   string `json:"propietario"`
		Datos         string `json:"datos"`
		Colaboradores string `json:"colaboradores"`
	}
	Calendarios []Calendario
)

//Crea una conexion con la DB y la retorna
func openDb() (db *sql.DB) {
	db, err := sql.Open("mysql", "root:etec@tcp(localhost:3306)/calendarios")
	if err != nil {
		panic("fallo la conexion con la base")
	}
	return db
}

//Busca y devuleve el id del ultimo calendario en toda la tabla
func lastIdCalendario(db *sql.DB) int {
	res, err := db.Query("SELECT idCalendario  FROM calendario ORDER BY idCalendario  DESC LIMIT 1")
	if err != nil {
		panic(err)
	}
	defer res.Close()

	var Id int

	for res.Next() {
		err = res.Scan(&Id)
		if err != nil {
			panic(err)
		}
	}
	return Id
}

//	RUTAS PARA USUARIOS

func (a *API) postUser(ec echo.Context) error {
	//	db := openDb()

	return ec.JSON(http.StatusAccepted, "Se guardo el nuevo usuario")
}

//	RUTAS PARA CALENDARIOS

//(id)-->(calendario) Devuelve un calenario
func (a *API) getCalendario(ec echo.Context) error {

	var c Calendario
	var cs Calendarios

	db := openDb()
	defer db.Close()

	id := ec.QueryParam("id")

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

//(user)-->([]calendarios) Devuleve todos los calendarios de un usuario
func (a *API) getCalendarios(ec echo.Context) error {

	var c Calendario
	var cs Calendarios

	db := openDb()
	defer db.Close()

	user := ec.QueryParam("propietario")

	rows, err := db.Query("SELECT * FROM calendario WHERE propietario = ?", user)
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

//(calendario)-->(ok) Crea un nuevo calendario
func (a *API) postCalendario(ec echo.Context) error {

	c := Calendario{
		Nombre:        ec.QueryParam("nombre"),
		Propietario:   ec.QueryParam("propietario"),
		Datos:         ec.QueryParam("datos"),
		Colaboradores: ec.QueryParam("colaboradores"),
	}

	db := openDb()
	defer db.Close()

	sentenciaPreparada, err := db.Prepare("INSERT INTO calendario (idCalendario, nombre, propietario, datos, colaboradores) VALUES(?, ?, ?, ?, ?)")
	if err != nil {
		log.Fatal(err)
	}
	defer sentenciaPreparada.Close()

	_, err = sentenciaPreparada.Exec(lastIdCalendario(db)+1, c.Nombre, c.Propietario, c.Datos, c.Colaboradores)

	if err != nil {
		log.Fatal(err)
	}

	return ec.JSON(http.StatusAccepted, "El usuario se registro con exito")
}

//(calendario, id)-->(ok) Cambia los datos de un calendario creado
func (a *API) patchCalendario(ec echo.Context) error {

	c := Calendario{
		Nombre:        ec.QueryParam("nombre"),
		Propietario:   ec.QueryParam("propietario"),
		Datos:         ec.QueryParam("datos"),
		Colaboradores: ec.QueryParam("colaboradores"),
	}

	var err error
	c.ID, err = strconv.Atoi(ec.QueryParam("id"))
	if err != nil {
		log.Fatal(err)
	}

	db := openDb()
	defer db.Close()

	sentenciaPreparada, err := db.Prepare("UPDATE calendario SET nombre = ?, propietario = ?, datos = ?, colaboradores = ? WHERE idCalendario = ?")
	if err != nil {
		log.Fatal(err)
	}
	defer sentenciaPreparada.Close()

	result, err := sentenciaPreparada.Exec(c.Nombre, c.Propietario, c.Datos, c.Colaboradores, c.ID)

	filasAfectadas, _ := result.RowsAffected()

	if filasAfectadas == 0 {
		return ec.JSON(http.StatusBadRequest, "no se encontro el valor que se quiere modificar")
	}

	if err != nil {
		log.Fatal("fallo la execucion de la centencia")
	}
	return ec.JSON(http.StatusAccepted, "se cambiaron los datos del calendario")
}

//(id)-->(ok) Elimina un calendario
func (a *API) deleteCalendario(ec echo.Context) error {
	id := ec.QueryParam("id")

	db := openDb()
	defer db.Close()

	sentenciaPreparada, err := db.Prepare("DELETE FROM calendario WHERE idCalendario = ?")
	if err != nil {
		log.Fatal(err)
	}
	defer sentenciaPreparada.Close()

	result, err := sentenciaPreparada.Exec(id)
	if err != nil {
		log.Fatal(err)
	}

	filasAfectadas, _ := result.RowsAffected()
	if filasAfectadas == 0 {
		return ec.JSON(http.StatusBadRequest, "no se encontro el valor que se quiere eliminar")
	}

	_, err = db.Query("UPDATE calendario SET idCalendario = ? WHERE idCalendario = ?", id, lastIdCalendario(db))
	if err != nil {
		log.Fatal(err)
	}

	return ec.JSON(http.StatusAccepted, "Se elimino el o los items de la tabla calendario")
}
