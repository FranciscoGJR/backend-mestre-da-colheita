package handlers

import (
   "github.com/gorilla/mux"
   "gorm.io/gorm"
)

func RegisterRoutes(r *mux.Router, db *gorm.DB) {
   r.HandleFunc("/api/v1/auth/register", RegisterUser(db)).Methods("POST")
   r.HandleFunc("/api/v1/auth/login", LoginUser(db)).Methods("POST")

   r.HandleFunc("/api/v1/culturas", ListCulturas(db)).Methods("GET")
   r.HandleFunc("/api/v1/culturas", CreateCultura(db)).Methods("POST")
   r.HandleFunc("/api/v1/culturas/{id}", UpdateCultura(db)).Methods("PUT")
   r.HandleFunc("/api/v1/culturas/{id}", DeleteCultura(db)).Methods("DELETE")

   r.HandleFunc("/api/v1/plantios", ListPlantios(db)).Methods("GET")
   r.HandleFunc("/api/v1/plantios", CreatePlantio(db)).Methods("POST")
   r.HandleFunc("/api/v1/plantios/{id}", GetPlantio(db)).Methods("GET")
   r.HandleFunc("/api/v1/plantios/{id}", UpdatePlantio(db)).Methods("PUT")
   r.HandleFunc("/api/v1/plantios/{id}", DeletePlantio(db)).Methods("DELETE")

   r.HandleFunc("/api/v1/colheitas", ListColheitas(db)).Methods("GET")
   r.HandleFunc("/api/v1/colheitas", CreateColheita(db)).Methods("POST")

   r.HandleFunc("/api/v1/dashboard", Dashboard(db)).Methods("GET")
   r.HandleFunc("/api/v1/relatorios/lucro", RelatorioLucro(db)).Methods("GET")
   r.HandleFunc("/api/v1/simulador", SimuladorLucro(db)).Methods("POST")
}
