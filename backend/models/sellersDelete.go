package models

import (
	"golang/db"
)

func SellersDelete(seller Sellers) (int64, error) {
	conn, err := db.OpenConnection()

	if err != nil {
		return 0, err
	}
	defer conn.Close()

	res, err := conn.Exec(`DELETE FROM sellers WHERE id=$1`, &seller.ID)
	if err != nil {
		return 0, err
	}

	return res.RowsAffected()
}
