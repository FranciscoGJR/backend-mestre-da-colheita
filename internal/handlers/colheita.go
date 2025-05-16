package handlers

import (
"encoding/json"
"net/http"
"time"

"gorm.io/gorm"
"github.com/FranciscoGJR/mestre-da-colheita/internal/models"
)

// GET /api/v1/colheitas
func ListColheitas(db *gorm.DB) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
    var colheitas []models.Colheita
    db.Preload("Itens").Find(&colheitas)
    w.Header().Set("Content-Type", "application/json")
    json.NewEncoder(w).Encode(colheitas)
	}
}

// POST /api/v1/colheitas
func CreateColheita(db *gorm.DB) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
    var req struct {
        IdPlantio    uint                   `json:"id_plantio"`
        DataColheita string                 `json:"data_colheita"`
        Itens        []models.ColheitaItem  `json:"itens"`
        Observacoes  string                 `json:"observacoes"`
    }
    if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
        http.Error(w, "JSON inválido", http.StatusBadRequest)
        return
    }
    dataColheita, _ := time.Parse("2006-01-02", req.DataColheita)
    colheita := models.Colheita{
        PlantioID:   req.IdPlantio,
        DataColheita: dataColheita,
        Observacoes: req.Observacoes,
        Itens:       req.Itens,
    }
    db.Create(&colheita)
    w.Header().Set("Content-Type", "application/json")
    w.WriteHeader(http.StatusCreated)
    json.NewEncoder(w).Encode(colheita)
	}
}
