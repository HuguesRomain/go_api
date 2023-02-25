package controllers

import (
	"encoding/json"
	"net/http"

	"github.com/HuguesRomain/go_api/internal/db"
	"github.com/HuguesRomain/go_api/models"
)

var shops []models.Shop


func CreateShop(w http.ResponseWriter, r *http.Request) {
	db, err := db.ConnectToDB()
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	defer db.Close()

	var shop models.Shop
	_ = json.NewDecoder(r.Body).Decode(&shop)

	// Créer le magasin dans la base de données
	if err := db.Create(&shop).Error; err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	json.NewEncoder(w).Encode(shop)
}


func GetShops(w http.ResponseWriter, r *http.Request) {
    w.Header().Set("Content-Type", "application/json")

    json.NewEncoder(w).Encode(shops)
}
