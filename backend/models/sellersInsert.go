package models

import "golang/db"

func SellersInsert(seller Sellers) (id int64, err error) {
	conn, err := db.OpenConnection()
	if err != nil {
		return
	}

	defer conn.Close()

	sql := `INSERT INTO sellers (name, email, document, address, created_at, updated_at) VALUES ($1, $2, $3, $4, $5, $6) RETURNING id`

	err = conn.QueryRow(sql, &seller.Name, &seller.Email, &seller.Document, &seller.Address, &seller.CreatedAt, &seller.UpdatedAt).Scan(&id)

	return
}
