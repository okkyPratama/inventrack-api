package repository

import (
	"database/sql"
	"inventrack/structs"
)

func GetAllProducts(db *sql.DB) ([]structs.Product, error) {
	query := `SELECT * FROM products`
	rows, err := db.Query(query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var products []structs.Product
	for rows.Next() {
		var p structs.Product
		err := rows.Scan(&p.ID, &p.Name, &p.Description, &p.Price, &p.Stock, &p.CategoryID, &p.SupplierID)
		if err != nil {
			return nil, err
		}
		products = append(products, p)
	}
	return products, nil
}

func GetProductByID(db *sql.DB, id int) (structs.Product, error) {
	query := `SELECT * FROM products WHERE id = $1`
	var p structs.Product
	err := db.QueryRow(query, id).Scan(&p.ID, &p.Name, &p.Description, &p.Price, &p.Stock, &p.CategoryID, &p.SupplierID)
	return p, err
}

func InsertProduct(db *sql.DB, p structs.Product) (int, error) {
	query := `INSERT INTO products (name, description, price, stock, category_id, supplier_id) VALUES ($1, $2, $3, $4, $5, $6) RETURNING id`
	var id int
	err := db.QueryRow(query, p.Name, p.Description, p.Price, p.Stock, p.CategoryID, p.SupplierID).Scan(&id)
	return id, err
}

func UpdateProduct(db *sql.DB, p structs.Product) error {
	query := `UPDATE products SET name = $1, description = $2, price = $3, stock = $4, category_id = $5, supplier_id = $6 WHERE id = $7`
	_, err := db.Exec(query, p.Name, p.Description, p.Price, p.Stock, p.CategoryID, p.SupplierID, p.ID)
	return err
}

func DeleteProduct(db *sql.DB, id int) error {
	query := `DELETE FROM products WHERE id = $1`
	_, err := db.Exec(query, id)
	return err
}
