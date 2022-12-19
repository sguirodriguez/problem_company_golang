package models

import (
	"golang/db"
)

func ProductsUpdate(product Products) (int64, error) {
	conn, err := db.OpenConnection()
	if err != nil {
		return 0, err
	}
	defer conn.Close()

	res, err := conn.Exec(`UPDATE products SET name=$2, value=$3, id_seller=$4 WHERE id=$1`, &product.ID, &product.Name, &product.Value, &product.IdSeller)
	if err != nil {
		return 0, err
	}

	return res.RowsAffected()
}
