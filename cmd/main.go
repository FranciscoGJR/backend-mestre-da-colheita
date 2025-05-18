package main

import (
	"log"
	"net/http"
	"os"

	"github.com/FranciscoGJR/mestre-da-colheita/internal/db"
	"github.com/FranciscoGJR/mestre-da-colheita/internal/handlers"
	"github.com/gorilla/mux"
	"github.com/joho/godotenv"
	httpSwagger "github.com/swaggo/http-swagger"
)

func enableCORS(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Access-Control-Allow-Origin", "*")
		w.Header().Set("Access-Control-Allow-Methods", "GET, POST, PUT, DELETE, OPTIONS")
		w.Header().Set("Access-Control-Allow-Headers", "Content-Type, Authorization")
		if r.Method == "OPTIONS" {
			w.WriteHeader(http.StatusOK)
			return
		}
		next.ServeHTTP(w, r)
	})
}

func main() {
	godotenv.Load()
	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}

	database := db.InitDB()
	db.SeedDatabase()

	r := mux.NewRouter()

	r.PathPrefix("/docs/").Handler(httpSwagger.WrapHandler)

	handlers.RegisterRoutes(r, database)

	// Adiciona o middleware de CORS
	handlerWithCORS := enableCORS(r)

	log.Printf("Servidor rodando na porta %s", port)
	log.Fatal(http.ListenAndServe(":"+port, handlerWithCORS))

}
