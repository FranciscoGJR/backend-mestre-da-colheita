package db

import (
	"fmt"
	"log"
	"os"
	"time"

	"github.com/FranciscoGJR/mestre-da-colheita/internal/models"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func SeedDatabase() {
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
		log.Fatal("Erro ao conectar ao banco:", err)
	}

	db.AutoMigrate(&models.Usuario{}, &models.Cultura{}, &models.Plantio{}, &models.Colheita{}, &models.ColheitaItem{})

	usuario := models.Usuario{
		Nome:  "Usuário Padrão",
		Email: "padrao@example.com",
		Senha: "123456",
	}
	db.FirstOrCreate(&usuario, models.Usuario{Email: usuario.Email})

	culturas := []models.Cultura{
		{
			Nome: "Tomate", TempoCrescimento: 10, PrecoNormal: 100,
			PrecoPrata: 120, PrecoOuro: 150, PrecoIridio: 180,
			Recorrente: true, ProdutividadeEsperada: 5, Estacao: "Verão",
		},
		{
			Nome: "Cenoura", TempoCrescimento: 7, PrecoNormal: 80,
			PrecoPrata: 100, PrecoOuro: 130, PrecoIridio: 160,
			Recorrente: false, ProdutividadeEsperada: 6, Estacao: "Primavera",
		},
	}
	for _, cultura := range culturas {
		db.FirstOrCreate(&cultura, models.Cultura{Nome: cultura.Nome})
	}

	var culturaTomate models.Cultura
	db.First(&culturaTomate, "nome = ?", "Tomate")

	plantio := models.Plantio{
		UsuarioID:          usuario.ID,
		CulturaID:          culturaTomate.ID,
		QuantidadeSementes: 100,
		DataPlantio:        time.Now().AddDate(0, 0, -10),
		Localizacao:        "Lote A",
		Status:             "Pronto para colher",
		Observacoes:        "Irrigação automática",
		FotoURL:            "http://example.com/foto.jpg",
	}
	db.Create(&plantio)

	colheita := models.Colheita{
		PlantioID:    plantio.ID,
		DataColheita: time.Now(),
		Observacoes:  "Colheita feita com sucesso",
		Itens: []models.ColheitaItem{
			{Qualidade: "Normal", Quantidade: 300},
			{Qualidade: "Ouro", Quantidade: 100},
		},
	}
	db.Create(&colheita)

	fmt.Println("🌱 Banco populado com sucesso!")
}
