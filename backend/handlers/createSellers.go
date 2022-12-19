package handlers

import (
	"encoding/json"
	"fmt"
	"golang/models"
	"log"
	"net/http"
)

func CreateSellers(w http.ResponseWriter, r *http.Request) {
	var seller models.Sellers

	err := json.NewDecoder(r.Body).Decode(&seller)
	if err != nil {
		log.Println("Erro ao fazer decode do json: %v", err)
		http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
		return
	}

	id, err := models.SellersInsert(seller)
	var resp map[string]any

	if err != nil {
		resp = map[string]any{
			"error": fmt.Sprintf("Ocorreu um erro ao tentar criar o vendedor: %v", err),
			"data":  false,
		}
	}
	if err == nil {
		resp = map[string]any{
			"error": false,
			"data":  fmt.Sprintf("Vendedor inserido com sucesso! ID: %d", id),
		}
	}
	w.Header().Add("Content-type", "application/json")
	json.NewEncoder(w).Encode(resp)

}
