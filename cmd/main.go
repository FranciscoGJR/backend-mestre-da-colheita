package main

import (
   "log"
   "net/http"
   "os"

   "github.com/gorilla/mux"
   "github.com/joho/godotenv"
   "github.com/swaggo/http-swagger"
   "github.com/FranciscoGJR/mestre-da-colheita/internal/db"
   "github.com/FranciscoGJR/mestre-da-colheita/internal/handlers"
)

func main() {
   godotenv.Load()
   port := os.Getenv("PORT")
   if port == "" {
       port = "8080"
   }

   database := db.InitDB(os.Getenv("DB_PATH"))
   r := mux.NewRouter()

   r.PathPrefix("/docs/").Handler(httpSwagger.WrapHandler)

   handlers.RegisterRoutes(r, database)

   log.Printf("Servidor rodando na porta %s", port)
   log.Fatal(http.ListenAndServe(":"+port, r))
}
