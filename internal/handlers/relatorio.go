package handlers

import (
"encoding/json"
"net/http"
"gorm.io/gorm"
"github.com/FranciscoGJR/mestre-da-colheita/internal/models"
)

// GET /api/v1/relatorios/lucro
func RelatorioLucro(db *gorm.DB) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
    var colheitas []models.Colheita
    db.Preload("Itens").Find(&colheitas)
    totalLucro := 0
    for _, c := range colheitas {
        for _, item := range c.Itens {
            preco := 100
            totalLucro += item.Quantidade * preco
        }
    }
    resp := map[string]interface{}{
        "lucro_total": totalLucro,
    }
    w.Header().Set("Content-Type", "application/json")
    json.NewEncoder(w).Encode(resp)
	}
}
