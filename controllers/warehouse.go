package controllers

import (
	"inventrack/database"
	"inventrack/repository"
	"inventrack/structs"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

func GetAllWarehouses(w http.ResponseWriter, r *http.Request) {
	warehouses, err := repository.GetAllWarehouses(database.DbConnection)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	respondJSON(w, http.StatusOK, warehouses)
}

func GetWarehouse(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id, err := strconv.Atoi(vars["id"])
	if err != nil {
		http.Error(w, "Invalid warehouse ID", http.StatusBadRequest)
		return
	}

	warehouse, err := repository.GetWarehouseByID(database.DbConnection, id)
	if err != nil {
		http.Error(w, "Warehouse not found", http.StatusNotFound)
		return
	}
	respondJSON(w, http.StatusOK, warehouse)
}

func CreateWarehouse(w http.ResponseWriter, r *http.Request) {
	var warehouse structs.Warehouse
	err := decodeJSONBody(r, &warehouse)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	id, err := repository.InsertWarehouse(database.DbConnection, warehouse)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	warehouse.ID = id
	respondJSON(w, http.StatusCreated, warehouse)
}

func UpdateWarehouse(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id, err := strconv.Atoi(vars["id"])
	if err != nil {
		http.Error(w, "Invalid warehouse ID", http.StatusBadRequest)
		return
	}

	var warehouse structs.Warehouse
	err = decodeJSONBody(r, &warehouse)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	warehouse.ID = id

	err = repository.UpdateWarehouse(database.DbConnection, warehouse)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	respondJSON(w, http.StatusOK, warehouse)
}

func DeleteWarehouse(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id, err := strconv.Atoi(vars["id"])
	if err != nil {
		http.Error(w, "Invalid warehouse ID", http.StatusBadRequest)
		return
	}

	err = repository.DeleteWarehouse(database.DbConnection, id)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	w.WriteHeader(http.StatusNoContent)
}
