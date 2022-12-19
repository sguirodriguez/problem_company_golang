package models

import "golang/db"

func ProductsInsert(product Products) (id int64, err error) {
	conn, err := db.OpenConnection()
	if err != nil {
		return
	}

	defer conn.Close()

	sql := `INSERT INTO products (name, value, id_seller, created_at, updated_at) VALUES ($1, $2, $3, $4, $5) RETURNING id`

	err = conn.QueryRow(sql, &product.Name, &product.Value, &product.IdSeller, &product.CreatedAt, &product.UpdatedAt).Scan(&id)

	return
}
