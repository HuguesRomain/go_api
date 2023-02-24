package main

import (
	"log"
	"net/http"

	"github.com/HuguesRomain/go_api/controllers"
	"github.com/gorilla/mux"
)

func main() {
    router := mux.NewRouter()

    router.HandleFunc("/shops", controllers.CreateShop).Methods("POST")
    router.HandleFunc("/shops", controllers.GetShops).Methods("GET")

    // Démarrage du serveur HTTP
    log.Fatal(http.ListenAndServe(":8000", router))
}
