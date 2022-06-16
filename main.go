package main

import (
	//Para imprimir en consola
	// El paquete HTTP

	// El paquete de rutas
	"github.com/joaco-basile/CFP/api" // El paquete de mi API
	"github.com/labstack/echo"
)

func main() {
	e := echo.New()
	//Creando un objeto API
	a := &api.API{}
	//Registrando las rutas
	a.RegisterRoutes(e)

	e.Logger.Fatal(e.Start(":8080"))

}
