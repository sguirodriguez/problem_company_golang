package models

import (
	"golang/db"
)

func DeliveriesGetAll() (deliveries []Deliveries, err error) {
	conn, err := db.OpenConnection()
	if err != nil {
		return
	}

	defer conn.Close()

	rows, err := conn.Query(`SELECT * FROM deliveries`)
	if err != nil {
		return
	}

	for rows.Next() {
		var delivery Deliveries

		err = rows.Scan(&delivery.ID, &delivery.ProductName, &delivery.Origin, &delivery.Destiny, &delivery.IdSeller, &delivery.IdClient, &delivery.IdProduct, &delivery.CreatedAt, &delivery.UpdatedAt, &delivery.ClientName)

		if err != nil {
			continue
		}

		deliveries = append(deliveries, delivery)
	}

	return
}
