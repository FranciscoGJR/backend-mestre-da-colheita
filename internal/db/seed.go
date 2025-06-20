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
			Nome: "Alho", TempoCrescimento: 4, PrecoNormal: 60, PrecoPrata: 75, PrecoOuro: 90, PrecoIridio: 120,
			Recorrente: false, ProdutividadeEsperada: 1, Estacao: "Primavera",
		},
		{
			Nome: "Batata", TempoCrescimento: 6, PrecoNormal: 80, PrecoPrata: 100, PrecoOuro: 120, PrecoIridio: 160,
			Recorrente: false, ProdutividadeEsperada: 1, Estacao: "Primavera",
		},
		{
			Nome: "Cenoura", TempoCrescimento: 3, PrecoNormal: 35, PrecoPrata: 43, PrecoOuro: 52, PrecoIridio: 70,
			Recorrente: false, ProdutividadeEsperada: 1, Estacao: "Primavera",
		},
		{
			Nome: "Chirívia", TempoCrescimento: 4, PrecoNormal: 35, PrecoPrata: 43, PrecoOuro: 52, PrecoIridio: 70,
			Recorrente: false, ProdutividadeEsperada: 1, Estacao: "Primavera",
		},
		{
			Nome: "Couve", TempoCrescimento: 6, PrecoNormal: 110, PrecoPrata: 137, PrecoOuro: 165, PrecoIridio: 220,
			Recorrente: false, ProdutividadeEsperada: 1, Estacao: "Primavera",
		},
		{
			Nome: "Morango", TempoCrescimento: 8, PrecoNormal: 120, PrecoPrata: 150, PrecoOuro: 180, PrecoIridio: 240,
			Recorrente: true, ProdutividadeEsperada: 1, Estacao: "Primavera",
		},
		{
			Nome: "Milho", TempoCrescimento: 14, PrecoNormal: 50, PrecoPrata: 62, PrecoOuro: 75, PrecoIridio: 100,
			Recorrente: true, ProdutividadeEsperada: 1, Estacao: "Verão",
		},
		{
			Nome: "Tomate", TempoCrescimento: 11, PrecoNormal: 60, PrecoPrata: 75, PrecoOuro: 90, PrecoIridio: 120,
			Recorrente: true, ProdutividadeEsperada: 1, Estacao: "Verão",
		},
		{
			Nome: "Trigo", TempoCrescimento: 4, PrecoNormal: 25, PrecoPrata: 31, PrecoOuro: 37, PrecoIridio: 50,
			Recorrente: false, ProdutividadeEsperada: 1, Estacao: "Outono",
		},
		{
			Nome: "Brócolis", TempoCrescimento: 8, PrecoNormal: 70, PrecoPrata: 87, PrecoOuro: 105, PrecoIridio: 140,
			Recorrente: true, ProdutividadeEsperada: 1, Estacao: "Outono",
		},
		{
			Nome: "Uva", TempoCrescimento: 10, PrecoNormal: 80, PrecoPrata: 100, PrecoOuro: 120, PrecoIridio: 160,
			Recorrente: true, ProdutividadeEsperada: 1, Estacao: "Outono",
		},
		{
			Nome: "Beterraba", TempoCrescimento: 6, PrecoNormal: 100, PrecoPrata: 125, PrecoOuro: 150, PrecoIridio: 200,
			Recorrente: false, ProdutividadeEsperada: 1, Estacao: "Outono",
		},
		{
			Nome: "Abacaxi", TempoCrescimento: 14, PrecoNormal: 300, PrecoPrata: 375, PrecoOuro: 450, PrecoIridio: 600,
			Recorrente: true, ProdutividadeEsperada: 1, Estacao: "Verão",
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
		Status:             "",
		Observacoes:        "Irrigação automática",
		FotoURL:            "http://example.com/foto.jpg",
	}
	db.Create(&plantio)

	colheita := models.Colheita{
		PlantioID:    plantio.ID,
		DataColheita: time.Now().AddDate(0, -2, 0),
		Observacoes:  "Colheita feita com sucesso",
		Itens: []models.ColheitaItem{
			{Qualidade: "Normal", Quantidade: 300},
			{Qualidade: "Ouro", Quantidade: 100},
		},
	}
	db.Create(&colheita)

	var culturaCenoura models.Cultura
	db.First(&culturaCenoura, "nome = ?", "Cenoura")

	plantioCenoura := models.Plantio{
		UsuarioID:          usuario.ID,
		CulturaID:          culturaCenoura.ID,
		QuantidadeSementes: 80,
		DataPlantio:        time.Now().AddDate(0, 0, -15), // 15 dias atrás
		Localizacao:        "Lote B",
		Status:             "",
		Observacoes:        "Plantio manual",
		FotoURL:            "http://example.com/foto-cenoura.jpg",
	}
	db.Create(&plantioCenoura)

	fmt.Println("🌱 Banco populado com sucesso!")
}
