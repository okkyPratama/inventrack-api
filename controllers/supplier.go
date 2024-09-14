package controllers

import (
	"inventrack/database"
	"inventrack/repository"
	"inventrack/structs"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

func GetAllSuppliers(w http.ResponseWriter, r *http.Request) {
	suppliers, err := repository.GetAllSuppliers(database.DbConnection)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	respondJSON(w, http.StatusOK, suppliers)
}

func GetSupplier(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id, err := strconv.Atoi(vars["id"])
	if err != nil {
		http.Error(w, "Invalid supplier ID", http.StatusBadRequest)
		return
	}

	supplier, err := repository.GetSupplierByID(database.DbConnection, id)
	if err != nil {
		http.Error(w, "Supplier not found", http.StatusNotFound)
		return
	}
	respondJSON(w, http.StatusOK, supplier)
}

func CreateSupplier(w http.ResponseWriter, r *http.Request) {
	var supplier structs.Supplier
	err := decodeJSONBody(r, &supplier)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	id, err := repository.InsertSupplier(database.DbConnection, supplier)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	supplier.ID = id
	respondJSON(w, http.StatusCreated, supplier)
}

func UpdateSupplier(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id, err := strconv.Atoi(vars["id"])
	if err != nil {
		http.Error(w, "Invalid supplier ID", http.StatusBadRequest)
		return
	}

	var supplier structs.Supplier
	err = decodeJSONBody(r, &supplier)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	supplier.ID = id

	err = repository.UpdateSupplier(database.DbConnection, supplier)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	respondJSON(w, http.StatusOK, supplier)
}

func DeleteSupplier(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id, err := strconv.Atoi(vars["id"])
	if err != nil {
		http.Error(w, "Invalid supplier ID", http.StatusBadRequest)
		return
	}

	err = repository.DeleteSupplier(database.DbConnection, id)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	w.WriteHeader(http.StatusNoContent)
}
