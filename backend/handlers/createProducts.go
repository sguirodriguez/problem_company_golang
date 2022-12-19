package handlers

import (
	"encoding/json"
	"fmt"
	"golang/models"
	"log"
	"net/http"
)

func CreateProducts(w http.ResponseWriter, r *http.Request) {
	var product models.Products

	err := json.NewDecoder(r.Body).Decode(&product)
	if err != nil {
		log.Println("Erro ao fazer decode do json: %v", err)
		http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
		return
	}

	id, err := models.ProductsInsert(product)
	var resp map[string]any

	if err != nil {
		resp = map[string]any{
			"error": fmt.Sprintf("Ocorreu um erro ao tentar criar o produto: %v", err),
			"data":  false,
		}
	}
	if err == nil {
		resp = map[string]any{
			"error": false,
			"data":  fmt.Sprintf("Produto inserido com sucesso! ID: %d", id),
		}
	}
	w.Header().Add("Content-type", "application/json")
	json.NewEncoder(w).Encode(resp)

}
