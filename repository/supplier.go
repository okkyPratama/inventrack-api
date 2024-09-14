package repository

import (
	"database/sql"
	"inventrack/structs"
)

func GetAllSuppliers(db *sql.DB) ([]structs.Supplier, error) {
	sql := "SELECT * FROM suppliers"
	rows, err := db.Query(sql)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var suppliers []structs.Supplier
	for rows.Next() {
		var supplier structs.Supplier
		err := rows.Scan(&supplier.ID, &supplier.Name, &supplier.ContactPerson, &supplier.Email, &supplier.Phone, &supplier.Address)
		if err != nil {
			return nil, err
		}
		suppliers = append(suppliers, supplier)
	}
	return suppliers, nil
}

func GetSupplierByID(db *sql.DB, id int) (structs.Supplier, error) {
	sql := "SELECT * FROM suppliers WHERE id = $1"
	var supplier structs.Supplier
	err := db.QueryRow(sql, id).Scan(&supplier.ID, &supplier.Name, &supplier.ContactPerson, &supplier.Email, &supplier.Phone, &supplier.Address)
	return supplier, err
}

func InsertSupplier(db *sql.DB, supplier structs.Supplier) (int, error) {
	sql := "INSERT INTO suppliers(name, contact_person, email, phone, address) VALUES ($1, $2, $3, $4, $5) RETURNING id"
	var id int
	err := db.QueryRow(sql, supplier.Name, supplier.ContactPerson, supplier.Email, supplier.Phone, supplier.Address).Scan(&id)
	return id, err
}

func UpdateSupplier(db *sql.DB, supplier structs.Supplier) error {
	sql := "UPDATE suppliers SET name = $1, contact_person = $2, email = $3, phone = $4, address = $5 WHERE id = $6"
	_, err := db.Exec(sql, supplier.Name, supplier.ContactPerson, supplier.Email, supplier.Phone, supplier.Address, supplier.ID)
	return err
}

func DeleteSupplier(db *sql.DB, id int) error {
	sql := "DELETE FROM suppliers WHERE id = $1"
	_, err := db.Exec(sql, id)
	return err
}
