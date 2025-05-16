package handlers

import (
"encoding/json"
"net/http"
"gorm.io/gorm"
"strconv"
"github.com/FranciscoGJR/mestre-da-colheita/internal/models"
)

// GET /api/v1/dashboard
func Dashboard(db *gorm.DB) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
    var plantios []models.Plantio
    db.Find(&plantios)

    alertas := []map[string]string{}
    for _, p := range plantios {
        if p.Status == "Pronto para colher" {
            alertas = append(alertas, map[string]string{
								"mensagem": "Amanhã: Cultura ID " + strconv.Itoa(int(p.CulturaID)) + " pronto para colher!",
                "tipo":     "colheita",
            })
        }
    }
    resp := map[string]interface{}{
        "plantios": plantios,
        "alertas":  alertas,
    }
    w.Header().Set("Content-Type", "application/json")
    json.NewEncoder(w).Encode(resp)
	}
}
