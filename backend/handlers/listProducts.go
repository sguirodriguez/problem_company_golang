package handlers

import (
	"encoding/json"
	"golang/models"
	"log"
	"net/http"
)

func ListProducts(w http.ResponseWriter, r *http.Request) {
	products, err := models.ProductsGetAll()

	if err != nil {
		log.Printf("Erro ao obter registros: %v", err)
	}

	w.Header().Add("Content-type", "application/json")
	w.WriteHeader(http.StatusCreated)

	var resp map[string]any

	resp = map[string]any{
		"error": false,
		"data":  products,
	}

	json.NewEncoder(w).Encode(resp)
}
