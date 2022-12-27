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
	// conexão db com docker
	conn := "host=localhost user=root password=root dbname=root port=5433 sslmode=disable"

	// conexão db local
	// conn := "host=localhost user=postgres password=123456 dbname=postgres port=5432 sslmode=disable"

	DB, err = gorm.Open(postgres.Open(conn))

	if err != nil {
		log.Panic("Erro na conexão com o banco de dados.")
	}

	DB.AutoMigrate(&models.Aluno{})
}
