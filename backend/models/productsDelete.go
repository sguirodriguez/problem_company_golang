package models

import (
	"golang/db"
)

func ProductsDelete(product Products) (int64, error) {
	conn, err := db.OpenConnection()

	if err != nil {
		return 0, err
	}
	defer conn.Close()

	res, err := conn.Exec(`DELETE FROM products WHERE id=$1`, &product.ID)
	if err != nil {
		return 0, err
	}

	return res.RowsAffected()
}
