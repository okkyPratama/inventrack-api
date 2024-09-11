package main

import (
	"database/sql"
	"fmt"
	"inventrack/auth"
	"inventrack/controllers"
	"inventrack/database"
	"inventrack/middleware"
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

func protectedHandler(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("This is a protected route"))
}

func main() {
	err = godotenv.Load("config/.env")
	if err != nil {
		panic("Error loading .env file")
	}

	psqlInfo := fmt.Sprintf(`host=%s port=%s user=%s password=%s dbname=%s sslmode=disable`,
		os.Getenv("DB_HOST"),
		os.Getenv("DB_PORT"),
		os.Getenv("DB_USER"),
		os.Getenv("DB_PASSWORD"),
		os.Getenv("DB_NAME"),
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
	r.HandleFunc("/api/protected", middleware.JWTMiddleware(protectedHandler)).Methods("GET")

	http.ListenAndServe(":8080", r)

}
