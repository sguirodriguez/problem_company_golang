package models

import (
	"golang/db"
)

func ProductsGetAll() (products []Products, err error) {
	conn, err := db.OpenConnection()
	if err != nil {
		return
	}

	defer conn.Close()

	rows, err := conn.Query(`SELECT * FROM products`)
	if err != nil {
		return
	}

	for rows.Next() {
		var product Products

		err = rows.Scan(&product.ID, &product.Name, &product.Value, &product.IdSeller, &product.CreatedAt, &product.UpdatedAt)

		if err != nil {
			continue
		}

		products = append(products, product)
	}

	return
}
