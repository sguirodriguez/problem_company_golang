package handlers

import (
	"encoding/json"
	"golang/models"

	"log"
	"net/http"
)

func UpdateSellers(w http.ResponseWriter, r *http.Request) {
	var seller models.Sellers

	var err = json.NewDecoder(r.Body).Decode(&seller)

	if err != nil {
		log.Println("Erro ao fazer decode do json: %v", err)
		http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
		return
	}

	rows, err := models.SellersUpdate(seller)
	if err != nil {
		log.Println("Erro ao atualizar registro: %v", err)
		http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
		return
	}

	if rows > 1 {
		log.Println("Error: foram atualizados %v registros", rows)
	}

	resp := map[string]any{
		"error": false,
		"data":  "dados atualizados com sucesso!",
	}
	w.Header().Add("Content-type", "application/json")
	json.NewEncoder(w).Encode(resp)
}
