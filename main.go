package main

import (
	"encoding/json"
	"log"
	"net/http" // El paquete HTTP
	"server/prueba/api"

	"github.com/gorilla/mux"
)

func calendario(w http.ResponseWriter, r *http.Request) {
	json.NewEncoder(w).Encode("{\"message\": \"hello word\"}")
}

func main() {
	router := mux.NewRouter()
	//Creando un objeto API
	a := &api.API{}
	//Registrando las rutas
	a.RegisterRoutes(router)

	router.HandleFunc("/", calendario).Methods(http.MethodGet)

	srv := &http.Server{
		Addr:    ":8080",
		Handler: router,
	}
	log.Println("Server en el puerto", srv.Addr)
	srv.ListenAndServe()
}
