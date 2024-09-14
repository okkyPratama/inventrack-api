package structs

import "time"

type Transaction struct {
	ID        int       `json:"id"`
	ProductID int       `json:"product_id"`
	Quantity  int       `json:"quantity"`
	Type      string    `json:"type"` // jika stok bertambah type nya "in", dan "out" jika stok nya berkurang
	Date      time.Time `json:"date"`
	UserID    int       `json:"user_id"`
}
