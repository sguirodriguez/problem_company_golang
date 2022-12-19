package models

import "golang/db"

func DeliveriesInsert(delivery Deliveries) (id int64, err error) {
	conn, err := db.OpenConnection()
	if err != nil {
		return
	}

	defer conn.Close()

	sql := `INSERT INTO deliveries (product_name, origin, destiny, id_seller, id_client, id_product, created_at, updated_at, client_name) VALUES ($1, $2, $3, $4, $5, $6, $7, $8, $9) RETURNING id`

	err = conn.QueryRow(sql, &delivery.ProductName, &delivery.Origin, &delivery.Destiny, &delivery.IdSeller, &delivery.IdClient, &delivery.IdProduct, &delivery.CreatedAt, &delivery.UpdatedAt, &delivery.ClientName).Scan(&id)

	return
}
