package db

import (
   "gorm.io/driver/sqlite"
   "gorm.io/gorm"
   "log"
   "github.com/FranciscoGJR/mestre-da-colheita/internal/models"
)

func InitDB(path string) *gorm.DB {
   db, err := gorm.Open(sqlite.Open(path), &gorm.Config{})
   if err != nil {
       log.Fatalf("Erro ao conectar no banco: %v", err)
   }
   // Auto-migrate models
   db.AutoMigrate(&models.Usuario{}, &models.Cultura{}, &models.Plantio{}, &models.Colheita{}, &models.ColheitaItem{})
   return db
}
