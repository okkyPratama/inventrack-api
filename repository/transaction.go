package repository

import (
	"database/sql"
	"inventrack/structs"
	"time"
)

type DBExecutor interface {
	Exec(query string, args ...interface{}) (sql.Result, error)
	Query(query string, args ...interface{}) (*sql.Rows, error)
	QueryRow(query string, args ...interface{}) *sql.Row
}

func CreateTransaction(db DBExecutor, t structs.Transaction) (int, error) {
	query := `INSERT INTO transactions (product_id, quantity, type, date, user_id) VALUES ($1, $2, $3, $4, $5) RETURNING id`
	var id int
	err := db.QueryRow(query, t.ProductID, t.Quantity, t.Type, time.Now(), t.UserID).Scan(&id)
	return id, err
}

func GetTransactionsByProductID(db DBExecutor, productID int) ([]structs.Transaction, error) {
	query := `SELECT id, product_id, quantity, type, date, user_id FROM transactions WHERE product_id = $1`
	rows, err := db.Query(query, productID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var transactions []structs.Transaction
	for rows.Next() {
		var t structs.Transaction
		err := rows.Scan(&t.ID, &t.ProductID, &t.Quantity, &t.Type, &t.Date, &t.UserID)
		if err != nil {
			return nil, err
		}
		transactions = append(transactions, t)
	}
	return transactions, nil
}

func UpdateProductStock(db DBExecutor, productID, quantity int) error {
	query := `UPDATE products SET stock = stock + $1 WHERE id = $2`
	_, err := db.Exec(query, quantity, productID)
	return err
}
