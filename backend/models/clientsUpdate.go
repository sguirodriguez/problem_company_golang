package models

import (
	"golang/db"
)

func ClientsUpdate(client Clients) (int64, error) {
	conn, err := db.OpenConnection()
	if err != nil {
		return 0, err
	}
	defer conn.Close()

	res, err := conn.Exec(`UPDATE clients SET name=$2, email=$3, document=$4, address=$5 WHERE id=$1`, &client.ID, &client.Name, &client.Email, &client.Document, &client.Address)
	if err != nil {
		return 0, err
	}

	return res.RowsAffected()
}
