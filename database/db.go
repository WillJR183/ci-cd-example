package database

import (
	"log"

	"github.com/WillJR183/ci-cd-example/models"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var (
	DB  *gorm.DB
	err error
)

func ConectaComBancoDeDados() {
	stringDeConexao := "host=localhost user=root password=root dbname=root port=5433 sslmode=disable"
	DB, err = gorm.Open(postgres.Open(stringDeConexao))

	if err != nil {
		log.Panic("Erro na conexão com o banco de dados.")
	}

	DB.AutoMigrate(&models.Aluno{})
}
