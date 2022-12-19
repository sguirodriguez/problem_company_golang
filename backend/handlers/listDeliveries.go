package handlers

import (
	"encoding/json"
	"golang/models"
	"log"
	"net/http"
)

func ListDeliveries(w http.ResponseWriter, r *http.Request) {
	deliveries, err := models.DeliveriesGetAll()

	if err != nil {
		log.Printf("Erro ao obter registros: %v", err)
	}

	w.Header().Add("Content-type", "application/json")
	w.WriteHeader(http.StatusCreated)

	var resp map[string]any

	resp = map[string]any{
		"error": false,
		"data":  deliveries,
	}

	json.NewEncoder(w).Encode(resp)
}
