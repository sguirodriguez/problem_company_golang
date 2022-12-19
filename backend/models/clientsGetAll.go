package models

import (
	"golang/db"
)

func ClientsGetAll() (clients []Clients, err error) {
	conn, err := db.OpenConnection()
	if err != nil {
		return
	}

	defer conn.Close()

	rows, err := conn.Query(`SELECT * FROM clients`)
	if err != nil {
		return
	}

	for rows.Next() {
		var client Clients

		err = rows.Scan(&client.ID, &client.Name, &client.Email, &client.Document, &client.Address, &client.CreatedAt, &client.UpdatedAt)

		if err != nil {
			continue
		}

		clients = append(clients, client)
	}

	return
}
