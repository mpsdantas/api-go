package main

import (
	"fmt"
	"log"
	"net/http"

	personController "./controllers/person"

	"github.com/gorilla/mux"
)

func main() {
	rotas := mux.NewRouter().StrictSlash(true)
	rotas.HandleFunc("/", personController.GetAll).Methods("GET")
	rotas.HandleFunc("/", personController.Create).Methods("POST")
	var port = ":3000"
	fmt.Println("Server running in port:", port)
	log.Fatal(http.ListenAndServe(port, rotas))
}
