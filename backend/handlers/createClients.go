package handlers

import (
	"encoding/json"
	"fmt"
	"golang/models"
	"log"
	"net/http"
)

func CreateClients(w http.ResponseWriter, r *http.Request) {
	var client models.Clients

	err := json.NewDecoder(r.Body).Decode(&client)
	if err != nil {
		log.Println("Erro ao fazer decode do json: %v", err)
		http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
		return
	}

	id, err := models.ClientsInsert(client)
	var resp map[string]any

	if err != nil {
		resp = map[string]any{
			"error": fmt.Sprintf("Ocorreu um erro ao tentar criar o cliente: %v", err),
			"data":  false,
		}
	}
	if err == nil {
		resp = map[string]any{
			"error": false,
			"data":  fmt.Sprintf("Cliente inserido com sucesso! ID: %d", id),
		}
	}
	w.Header().Add("Content-type", "application/json")
	json.NewEncoder(w).Encode(resp)

}
