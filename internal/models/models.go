package models

import (
   "gorm.io/gorm"
   "time"
)

type Usuario struct {
   ID       uint   `gorm:"primaryKey" json:"id_usuario"`
   Nome     string `json:"nome"`
   Email    string `gorm:"unique" json:"email"`
   Senha    string `json:"senha"`
}

type Cultura struct {
   ID                   uint   `gorm:"primaryKey" json:"id"`
   Nome                 string `gorm:"unique" json:"nome"`
   TempoCrescimento     int    `json:"tempo_crescimento"`
   PrecoNormal          int    `json:"preco_normal"`
   PrecoPrata           int    `json:"preco_prata"`
   PrecoOuro            int    `json:"preco_ouro"`
   PrecoIridio          int    `json:"preco_iridio"`
   Recorrente           bool   `json:"recorrente"`
   ProdutividadeEsperada int   `json:"produtividade_esperada"`
   PrecoCustomizavel    bool   `json:"preco_customizavel"`
}

type Plantio struct {
   ID                uint      `gorm:"primaryKey" json:"id"`
   UsuarioID         uint      `json:"id_usuario"`
   CulturaID         uint      `json:"id_cultura"`
   QuantidadeSementes int      `json:"quantidade_sementes"`
   DataPlantio       time.Time `json:"data_plantio"`
   EstacaoPlantio    string    `json:"estacao_plantio"`
   Localizacao       string    `json:"localizacao"`
   Status            string    `json:"status"`
   Observacoes       string    `json:"observacoes"`
   FotoURL           string    `json:"foto_url"`
}

type Colheita struct {
   ID         uint      `gorm:"primaryKey" json:"id"`
   PlantioID  uint      `json:"id_plantio"`
   DataColheita time.Time `json:"data_colheita"`
   Observacoes string    `json:"observacoes"`
   Itens      []ColheitaItem `gorm:"foreignKey:ColheitaID" json:"itens"`
}

type ColheitaItem struct {
   ID         uint   `gorm:"primaryKey" json:"id"`
   ColheitaID uint   `json:"colheita_id"`
   Qualidade  string `json:"qualidade"`
   Quantidade int    `json:"quantidade"`
}
