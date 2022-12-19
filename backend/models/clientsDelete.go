package models

import (
	"golang/db"
)

func ClientsDelete(client Clients) (int64, error) {
	conn, err := db.OpenConnection()

	if err != nil {
		return 0, err
	}
	defer conn.Close()

	res, err := conn.Exec(`DELETE FROM clients WHERE id=$1`, &client.ID)
	if err != nil {
		return 0, err
	}

	return res.RowsAffected()
}
