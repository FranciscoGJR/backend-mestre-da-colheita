package db

import (
"fmt"
"log"
"os"

"github.com/joho/godotenv"
"gorm.io/driver/postgres"
"gorm.io/gorm"
"github.com/FranciscoGJR/mestre-da-colheita/internal/models"
)

func InitDB() *gorm.DB {
	err := godotenv.Load()
	if err != nil {
    log.Println("Aviso: .env não encontrado, usando variáveis de ambiente do sistema")
	}

	host := os.Getenv("DB_HOST")
	user := os.Getenv("DB_USER")
	password := os.Getenv("DB_PASSWORD")
	dbname := os.Getenv("DB_NAME")
	port := os.Getenv("DB_PORT")
	sslmode := os.Getenv("DB_SSLMODE")

	dsn := fmt.Sprintf(
    "host=%s user=%s password=%s dbname=%s port=%s sslmode=%s",
    host, user, password, dbname, port, sslmode,
	)

	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
    log.Fatalf("Erro ao conectar no banco: %v", err)
	}
// Auto-migrate models
	db.AutoMigrate(&models.Usuario{}, &models.Cultura{}, &models.Plantio{}, &models.Colheita{}, &models.ColheitaItem{})
	return db
}
