package handlers

import (
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/FranciscoGJR/mestre-da-colheita/internal/models"
	"github.com/gorilla/mux"
	"gorm.io/gorm"
)

func ListCulturas(db *gorm.DB) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var culturas []models.Cultura
		db.Find(&culturas)
		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(culturas)
	}
}

func CreateCultura(db *gorm.DB) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var cultura models.Cultura
		if err := json.NewDecoder(r.Body).Decode(&cultura); err != nil {
			http.Error(w, "JSON inválido", http.StatusBadRequest)
			return
		}
		if cultura.Nome == "" || cultura.TempoCrescimento <= 0 {
			http.Error(w, "Dados obrigatórios faltando", http.StatusBadRequest)
			return
		}
		db.Create(&cultura)
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusCreated)
		json.NewEncoder(w).Encode(cultura)
	}
}

func UpdateCultura(db *gorm.DB) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		id, _ := strconv.Atoi(mux.Vars(r)["id"])
		var cultura models.Cultura
		if err := db.First(&cultura, id).Error; err != nil {
			http.Error(w, "Cultura não encontrada", http.StatusNotFound)
			return
		}
		var update models.Cultura
		if err := json.NewDecoder(r.Body).Decode(&update); err != nil {
			http.Error(w, "JSON inválido", http.StatusBadRequest)
			return
		}
		db.Model(&cultura).Updates(update)
		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(cultura)
	}
}

func DeleteCultura(db *gorm.DB) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		idStr := mux.Vars(r)["id"]
		id, err := strconv.Atoi(idStr)
		if err != nil {
			http.Error(w, "ID inválido", http.StatusBadRequest)
			return
		}
		db.Delete(&models.Cultura{}, id)
		w.WriteHeader(http.StatusNoContent)
	}
}

func GetCulturaByID(db *gorm.DB) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		idStr := mux.Vars(r)["id"]
		id, err := strconv.Atoi(idStr)
		if err != nil {
			http.Error(w, "ID inválido", http.StatusBadRequest)
			return
		}
		var cultura models.Cultura
		if err := db.First(&cultura, id).Error; err != nil {
			http.Error(w, "Cultura não encontrada", http.StatusNotFound)
			return
		}
		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(cultura)
	}
}
