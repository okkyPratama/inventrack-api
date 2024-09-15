package controllers

import (
	"database/sql"
	"encoding/json"
	"inventrack/database"
	"inventrack/repository"
	"inventrack/structs"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

func GetAllProducts(w http.ResponseWriter, r *http.Request) {
	products, err := repository.GetAllProducts(database.DbConnection)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	json.NewEncoder(w).Encode(products)
}

func GetProduct(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id, err := strconv.Atoi(vars["id"])
	if err != nil {
		http.Error(w, "Invalid product ID", http.StatusBadRequest)
		return
	}

	product, err := repository.GetProductByID(database.DbConnection, id)
	if err != nil {
		http.Error(w, "Product not found", http.StatusNotFound)
		return
	}
	json.NewEncoder(w).Encode(product)
}

func CreateProduct(w http.ResponseWriter, r *http.Request) {
	var product structs.Product
	err := json.NewDecoder(r.Body).Decode(&product)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	id, err := repository.InsertProduct(database.DbConnection, product)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	product.ID = id
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(product)
}

func UpdateProduct(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id, err := strconv.Atoi(vars["id"])
	if err != nil {
		http.Error(w, "Invalid product ID", http.StatusBadRequest)
		return
	}

	var product structs.Product
	err = json.NewDecoder(r.Body).Decode(&product)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	product.ID = id

	err = repository.UpdateProduct(database.DbConnection, product)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	json.NewEncoder(w).Encode(product)
}

func DeleteProduct(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id, err := strconv.Atoi(vars["id"])
	if err != nil {
		respondJSON(w, http.StatusBadRequest, map[string]string{"error": "Invalid product ID"})
		return
	}

	err = repository.DeleteProduct(database.DbConnection, id)
	if err != nil {
		if err == sql.ErrNoRows {
			respondJSON(w, http.StatusNotFound, map[string]string{"error": "Product not found"})
		} else {
			respondJSON(w, http.StatusInternalServerError, map[string]string{"error": "Failed to delete product"})
		}
		return
	}
	respondJSON(w, http.StatusOK, map[string]string{"message": "Product deleted successfully"})
}
