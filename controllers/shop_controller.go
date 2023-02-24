// Dans le fichier shop_controller.go
package controllers

import (
	"encoding/json"
	"net/http"

	"github.com/HuguesRomain/go_api/models"
)

var shops []models.Shop

func CreateShop(w http.ResponseWriter, r *http.Request) {
    w.Header().Set("Content-Type", "application/json")

    var shop models.Shop
    _ = json.NewDecoder(r.Body).Decode(&shop)

    shops = append(shops, shop)

    json.NewEncoder(w).Encode(shop)
}

func GetShops(w http.ResponseWriter, r *http.Request) {
    w.Header().Set("Content-Type", "application/json")

    json.NewEncoder(w).Encode(shops)
}
