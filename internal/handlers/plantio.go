package handlers

import (
"encoding/json"
"net/http"
"strconv"
"time"

"github.com/gorilla/mux"
"gorm.io/gorm"
"github.com/FranciscoGJR/mestre-da-colheita/internal/models"
)

// GET /api/v1/plantios
func ListPlantios(db *gorm.DB) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var plantios []models.Plantio
    db.Find(&plantios)
    w.Header().Set("Content-Type", "application/json")
    json.NewEncoder(w).Encode(plantios)
	}
}

// POST /api/v1/plantios
func CreatePlantio(db *gorm.DB) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req models.Plantio
    if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
        http.Error(w, "JSON inválido", http.StatusBadRequest)
        return
    }

    req.Status = "Crescendo"
    if req.DataPlantio.IsZero() {
        req.DataPlantio = time.Now()
    }
    db.Create(&req)
    w.Header().Set("Content-Type", "application/json")
    w.WriteHeader(http.StatusCreated)
    json.NewEncoder(w).Encode(req)
	}
}

// GET /api/v1/plantios/{id}
func GetPlantio(db *gorm.DB) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		id, _ := strconv.Atoi(mux.Vars(r)["id"])
    var plantio models.Plantio
    if err := db.First(&plantio, id).Error; err != nil {
        http.Error(w, "Plantio não encontrado.", http.StatusNotFound)
        return
    }
    w.Header().Set("Content-Type", "application/json")
    json.NewEncoder(w).Encode(plantio)
	}
}

// PUT /api/v1/plantios/{id}
func UpdatePlantio(db *gorm.DB) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
    id, _ := strconv.Atoi(mux.Vars(r)["id"])
    var plantio models.Plantio
    if err := db.First(&plantio, id).Error; err != nil {
        http.Error(w, "Plantio não encontrado.", http.StatusNotFound)
        return
    }
    var update models.Plantio
    if err := json.NewDecoder(r.Body).Decode(&update); err != nil {
        http.Error(w, "JSON inválido", http.StatusBadRequest)
        return
    }
    db.Model(&plantio).Updates(update)
    w.Header().Set("Content-Type", "application/json")
    json.NewEncoder(w).Encode(plantio)
	}
}

// DELETE /api/v1/plantios/{id}
func DeletePlantio(db *gorm.DB) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
    id, _ := strconv.Atoi(mux.Vars(r)["id"])
    db.Delete(&models.Plantio{}, id)
    w.WriteHeader(http.StatusNoContent)
	}
}
