package models

import "golang/db"

func ClientsInsert(client Clients) (id int64, err error) {
	conn, err := db.OpenConnection()
	if err != nil {
		return
	}

	defer conn.Close()

	sql := `INSERT INTO clients (name, email, document, address, created_at, updated_at) VALUES ($1, $2, $3, $4, $5, $6) RETURNING id`

	err = conn.QueryRow(sql, &client.Name, &client.Email, &client.Document, &client.Address, &client.CreatedAt, &client.UpdatedAt).Scan(&id)

	return
}
