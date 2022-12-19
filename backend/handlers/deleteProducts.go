package handlers

import (
	"encoding/json"
	"golang/models"
	"log"
	"net/http"
)

func DeleteProducts(w http.ResponseWriter, r *http.Request) {
	var product models.Products

	var err = json.NewDecoder(r.Body).Decode(&product)

	rows, err := models.ProductsDelete(product)
	if err != nil {
		log.Println("Erro ao deletar registro: %v", err)
		http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
		return
	}

	if rows > 1 {
		log.Println("Error: foram removidos %v registros", rows)
	}

	resp := map[string]any{
		"error": false,
		"data":  "dados removidos com sucesso!",
	}
	w.Header().Add("Content-type", "application/json")
	json.NewEncoder(w).Encode(resp)
}
