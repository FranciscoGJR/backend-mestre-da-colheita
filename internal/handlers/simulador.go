package handlers

import (
"encoding/json"
"net/http"
"gorm.io/gorm"
"github.com/FranciscoGJR/mestre-da-colheita/internal/models"
)

// POST /api/v1/simulador
func SimuladorLucro(db *gorm.DB) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
    var req struct {
        IDCultura uint `json:"id_cultura"`
        Quantidade int `json:"quantidade"`
    }
    if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
        http.Error(w, "JSON inválido", http.StatusBadRequest)
        return
    }
    var cultura models.Cultura
    if err := db.First(&cultura, req.IDCultura).Error; err != nil {
        http.Error(w, "Cultura não encontrada", http.StatusNotFound)
        return
    }
    lucroEstimado := req.Quantidade * cultura.PrecoNormal
    resp := map[string]interface{}{
        "lucro_estimado": lucroEstimado,
        "dias_ciclo":     cultura.TempoCrescimento,
    }
    w.Header().Set("Content-Type", "application/json")
    json.NewEncoder(w).Encode(resp)
	}
}
