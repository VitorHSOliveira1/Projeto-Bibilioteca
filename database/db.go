package database

import (
	"Projeto-da-Bibilioteca/models"
	"log"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var (
	DB  *gorm.DB
	err error
)

func ConectaComBancoDeDados() {
	stringDeConexao := "host=localhost user=root password=root dbname=root port=5432 sslmode=disable"
	DB, err = gorm.Open(postgres.Open(stringDeConexao), &gorm.Config{})
	if err != nil {
		log.Panic("Erro ao conectar com o banco de dados: ", err)
	}

	err = DB.AutoMigrate(&models.Livro{})
	if err != nil {
		log.Panic("Erro ao realizar AutoMigrate: ", err)
	}

	log.Println("Conexão com o banco de dados estabelecida com sucesso e AutoMigrate concluído!")
}
