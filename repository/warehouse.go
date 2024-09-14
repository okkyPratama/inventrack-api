package repository

import (
	"database/sql"
	"inventrack/structs"
)

func GetAllWarehouses(db *sql.DB) ([]structs.Warehouse, error) {
	sql := "SELECT * FROM warehouses"
	rows, err := db.Query(sql)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var warehouses []structs.Warehouse
	for rows.Next() {
		var warehouse structs.Warehouse
		err := rows.Scan(&warehouse.ID, &warehouse.Name, &warehouse.Description, &warehouse.Capacity)
		if err != nil {
			return nil, err
		}
		warehouses = append(warehouses, warehouse)
	}
	return warehouses, nil
}

func GetWarehouseByID(db *sql.DB, id int) (structs.Warehouse, error) {
	sql := "SELECT * FROM warehouses WHERE id = $1"
	var warehouse structs.Warehouse
	err := db.QueryRow(sql, id).Scan(&warehouse.ID, &warehouse.Name, &warehouse.Description, &warehouse.Capacity)
	return warehouse, err
}

func InsertWarehouse(db *sql.DB, warehouse structs.Warehouse) (int, error) {
	sql := "INSERT INTO warehouses(name, description, capacity) VALUES ($1, $2, $3) RETURNING id"
	var id int
	err := db.QueryRow(sql, warehouse.Name, warehouse.Description, warehouse.Capacity).Scan(&id)
	return id, err
}

func UpdateWarehouse(db *sql.DB, warehouse structs.Warehouse) error {
	sql := "UPDATE warehouses SET name = $1, description = $2, capacity = $3 WHERE id = $4"
	_, err := db.Exec(sql, warehouse.Name, warehouse.Description, warehouse.Capacity, warehouse.ID)
	return err
}

func DeleteWarehouse(db *sql.DB, id int) error {
	sql := "DELETE FROM warehouses WHERE id = $1"
	_, err := db.Exec(sql, id)
	return err
}
