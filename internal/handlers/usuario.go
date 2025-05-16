package handlers

import (
   "encoding/json"
   "net/http"
   "gorm.io/gorm"
   "github.com/seuusuario/plantio-smart/internal/models"
)

func RegisterUser(db *gorm.DB) http.HandlerFunc {
   return func(w http.ResponseWriter, r *http.Request) {
       var user models.Usuario
       if err := json.NewDecoder(r.Body).Decode(&user); err != nil {
           http.Error(w, "JSON inválido", http.StatusBadRequest)
           return
       }

       if user.Email == "" || len(user.Senha) < 6 {
           http.Error(w, "Email ou senha inválidos", http.StatusBadRequest)
           return
       }

       var exists models.Usuario
       if err := db.Where("email = ?", user.Email).First(&exists).Error; err == nil {
           http.Error(w, "Email já cadastrado.", http.StatusConflict)
           return
       }
       db.Create(&user)
       user.Senha = ""
       w.Header().Set("Content-Type", "application/json")
       w.WriteHeader(http.StatusCreated)
       json.NewEncoder(w).Encode(user)
   }
}

func LoginUser(db *gorm.DB) http.HandlerFunc {
   return func(w http.ResponseWriter, r *http.Request) {
       var req struct {
           Email string `json:"email"`
           Senha string `json:"senha"`
       }
       if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
           http.Error(w, "JSON inválido", http.StatusBadRequest)
           return
       }
       var user models.Usuario
       if err := db.Where("email = ? AND senha = ?", req.Email, req.Senha).First(&user).Error; err != nil {
           http.Error(w, "Usuário ou senha inválidos", http.StatusUnauthorized)
           return
       }
       user.Senha = ""
       w.Header().Set("Content-Type", "application/json")
       json.NewEncoder(w).Encode(user)
   }
}
