package main

import (
	"database/sql"
	"fmt"
	"inventrack/auth"
	"inventrack/controllers"
	"inventrack/database"
	"inventrack/middleware"
	"log"
	"net/http"
	"os"

	"github.com/gorilla/mux"
	"github.com/joho/godotenv"
	_ "github.com/lib/pq"
)

var (
	DB  *sql.DB
	err error
)

func main() {
	err = godotenv.Load("config/.env")
	if err != nil {
		panic("Error loading .env file")
	}

	psqlInfo := fmt.Sprintf(`host=%s port=%s user=%s password=%s dbname=%s sslmode=disable`,
		// local config
		// os.Getenv("DB_HOST"),
		// os.Getenv("DB_PORT"),
		// os.Getenv("DB_USER"),
		// os.Getenv("DB_PASSWORD"),
		// os.Getenv("DB_NAME"),

		// prod config
		os.Getenv("PGHOST"),
		os.Getenv("PGPORT"),
		os.Getenv("PGUSER"),
		os.Getenv("PGPASSWORD"),
		os.Getenv("PGDATABASE"),
	)

	DB, err = sql.Open("postgres", psqlInfo)
	defer DB.Close()
	err = DB.Ping()
	if err != nil {
		panic(err)
	}

	database.DBMigrate(DB)

	fmt.Println("Successfully connected!")

	jwtSecret := os.Getenv("JWT_SECRET_KEY")
	if jwtSecret == "" {
		panic("JWT_SECRET_KEY not found")
	}
	auth.SetJWTSecretKey(jwtSecret)

	r := mux.NewRouter()

	// Public routes
	r.HandleFunc("/api/users/register", controllers.Register).Methods("POST")
	r.HandleFunc("/api/users/login", controllers.Login).Methods("POST")

	//Protected routes
	api := r.PathPrefix("/api").Subrouter()
	api.Use(middleware.JWTMiddleware)

	// Category routes
	api.HandleFunc("/categories", controllers.GetAllCategories).Methods("GET")
	api.HandleFunc("/categories", controllers.CreateCategory).Methods("POST")
	api.HandleFunc("/categories/{id}", controllers.GetCategory).Methods("GET")
	api.HandleFunc("/categories/{id}", controllers.UpdateCategory).Methods("PUT")
	api.HandleFunc("/categories/{id}", controllers.DeleteCategory).Methods("DELETE")

	// Supplier routes
	api.HandleFunc("/suppliers", controllers.GetAllSuppliers).Methods("GET")
	api.HandleFunc("/suppliers", controllers.CreateSupplier).Methods("POST")
	api.HandleFunc("/suppliers/{id}", controllers.GetSupplier).Methods("GET")
	api.HandleFunc("/suppliers/{id}", controllers.UpdateSupplier).Methods("PUT")
	api.HandleFunc("/suppliers/{id}", controllers.DeleteSupplier).Methods("DELETE")

	// Product routes
	api.HandleFunc("/products", controllers.GetAllProducts).Methods("GET")
	api.HandleFunc("/products", controllers.CreateProduct).Methods("POST")
	api.HandleFunc("/products/{id}", controllers.GetProduct).Methods("GET")
	api.HandleFunc("/products/{id}", controllers.UpdateProduct).Methods("PUT")
	api.HandleFunc("/products/{id}", controllers.DeleteProduct).Methods("DELETE")

	// Transaction routes
	api.HandleFunc("/transactions", controllers.CreateTransaction).Methods("POST")
	api.HandleFunc("/products/{id}/transactions", controllers.GetProductTransactions).Methods("GET")

	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}
	log.Fatal(http.ListenAndServe(":"+port, r))
}
