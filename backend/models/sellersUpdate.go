package models

import (
	"golang/db"
)

func SellersUpdate(seller Sellers) (int64, error) {
	conn, err := db.OpenConnection()
	if err != nil {
		return 0, err
	}
	defer conn.Close()

	res, err := conn.Exec(`UPDATE sellers SET name=$2, email=$3, document=$4, address=$5 WHERE id=$1`, &seller.ID, &seller.Name, &seller.Email, &seller.Document, &seller.Address)
	if err != nil {
		return 0, err
	}

	return res.RowsAffected()
}
