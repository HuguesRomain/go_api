package main

import (
	"log"
	"net/http"

	"github.com/HuguesRomain/go_api/controllers"
	"github.com/HuguesRomain/go_api/internal/db"
	"github.com/gorilla/mux"
	_ "github.com/jinzhu/gorm/dialects/postgres"
)


func main() {
    db, err := db.ConnectToDB()
	if err != nil {
		log.Fatalf("could not connect to database: %v", err)
	}
	defer db.Close()

    
    router := mux.NewRouter()

    router.HandleFunc("/shops", controllers.CreateShop).Methods("POST")
    router.HandleFunc("/shops", controllers.GetShops).Methods("GET")

    // Démarrage du serveur HTTP

    log.Fatal(http.ListenAndServe(":8000", router))
}
