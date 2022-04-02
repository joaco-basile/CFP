package main

import (
	"log"      //Para imprimir en consola
	"net/http" // El paquete HTTP

	"github.com/gorilla/mux"          // El paquete de rutas
	"github.com/joaco-basile/CFP/api" // El paquete de mi API
)

func main() {
	router := mux.NewRouter()
	//Creando un objeto API
	a := &api.API{}
	//Registrando las rutas
	a.RegisterRoutes(router)

	srv := &http.Server{
		Addr:    ":8080",
		Handler: router,
	}

	log.Println("Server en el puerto", srv.Addr)
	srv.ListenAndServe()
}
