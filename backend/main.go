package main

import (
	"fmt"
	"golang/configs"
	"golang/handlers"
	"net/http"

	"github.com/go-chi/chi/v5"
)

func main() {
	err := configs.Load()
	if err != nil {

		panic(err)

	}

	r := chi.NewRouter()

	r.Post("/seller", handlers.CreateSellers)
	r.Patch("/seller", handlers.UpdateSellers)
	r.Delete("/seller", handlers.DeleteSellers)
	r.Get("/seller", handlers.ListSellers)

	r.Post("/client", handlers.CreateClients)
	r.Patch("/client", handlers.UpdateClients)
	r.Delete("/client", handlers.DeleteClients)
	r.Get("/client", handlers.ListClients)

	r.Post("/product", handlers.CreateProducts)
	r.Patch("/product", handlers.UpdateProducts)
	r.Delete("/product", handlers.DeleteProducts)
	r.Get("/product", handlers.ListProducts)

	r.Get("/delivery", handlers.ListDeliveries)

	http.ListenAndServe(fmt.Sprintf(":%s", configs.GetServerPort()), r)
}
