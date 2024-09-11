package controllers

import (
	"encoding/json"
	"inventrack/auth"
	"inventrack/database"
	"inventrack/structs"
	"net/http"
)

func Register(w http.ResponseWriter, r *http.Request) {
	var user structs.User
	json.NewDecoder(r.Body).Decode(&user)

	hashedPassword, err := auth.HashPassword(user.Password)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	_, err = database.DbConnection.Exec("INSERT INTO users (username, email, password, role) VALUES ($1, $2, $3, $4)",
		user.Username, user.Email, hashedPassword, user.Role)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(map[string]string{"message": "User registered successfully"})
}

func Login(w http.ResponseWriter, r *http.Request) {
	var loginUser structs.User
	json.NewDecoder(r.Body).Decode(&loginUser)

	var user structs.User
	err := database.DbConnection.QueryRow("SELECT id, username, password FROM users WHERE username = $1", loginUser.Username).Scan(&user.ID, &user.Username, &user.Password)
	if err != nil {
		http.Error(w, "Invalid credentials", http.StatusUnauthorized)
		return
	}

	if !auth.CheckPasswordHash(loginUser.Password, user.Password) {
		http.Error(w, "Invalid credentials", http.StatusUnauthorized)
		return
	}

	token, err := auth.GenerateToken(user.ID)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	json.NewEncoder(w).Encode(map[string]string{"token": token})
}
