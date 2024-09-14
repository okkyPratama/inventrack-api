package repository

import (
	"database/sql"
	"inventrack/structs"
)

func GetAllCategories(db *sql.DB) ([]structs.Category, error) {
	sql := "SELECT * FROM categories"
	rows, err := db.Query(sql)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var categories []structs.Category
	for rows.Next() {
		var category structs.Category
		err := rows.Scan(&category.ID, &category.Name, &category.Description)
		if err != nil {
			return nil, err
		}
		categories = append(categories, category)
	}
	return categories, nil
}

func GetCategoryByID(db *sql.DB, id int) (structs.Category, error) {
	sql := "SELECT * FROM categories WHERE id = $1"
	var category structs.Category
	err := db.QueryRow(sql, id).Scan(&category.ID, &category.Name, &category.Description)
	return category, err
}

func InsertCategory(db *sql.DB, category structs.Category) (int, error) {
	sql := "INSERT INTO categories(name, description) VALUES ($1, $2) RETURNING id"
	var id int
	err := db.QueryRow(sql, category.Name, category.Description).Scan(&id)
	return id, err
}

func UpdateCategory(db *sql.DB, category structs.Category) error {
	sql := "UPDATE categories SET name = $1, description = $2 WHERE id = $3"
	_, err := db.Exec(sql, category.Name, category.Description, category.ID)
	return err
}

func DeleteCategory(db *sql.DB, id int) error {
	sql := "DELETE FROM categories WHERE id = $1"
	_, err := db.Exec(sql, id)
	return err
}
