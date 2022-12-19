package models

import (
	"golang/db"
)

func SellersGetAll() (sellers []Sellers, err error) {
	conn, err := db.OpenConnection()
	if err != nil {
		return
	}

	defer conn.Close()

	rows, err := conn.Query(`SELECT * FROM sellers`)
	if err != nil {
		return
	}

	for rows.Next() {
		var seller Sellers

		err = rows.Scan(&seller.ID, &seller.Name, &seller.Email, &seller.Document, &seller.Address, &seller.CreatedAt, &seller.UpdatedAt)

		if err != nil {
			continue
		}

		sellers = append(sellers, seller)
	}

	return
}
