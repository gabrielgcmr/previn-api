package database

import (
	"fmt"
	"log"
	"os"
	"time"

	"github.com/joho/godotenv"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var DB *gorm.DB

func Connect() {
	// Carrega o .env
	err := godotenv.Load()
	if err != nil {
		log.Println("⚠️  Arquivo .env não encontrado. Continuando com variáveis de ambiente do sistema...")
	}

	host := os.Getenv("_DB_HOST")
	port := os.Getenv("_DB_PORT")
	user := os.Getenv("_DB_USER")
	password := os.Getenv("_DB_PASSWORD")
	dbname := os.Getenv("_DB_NAME")

	dsn := fmt.Sprintf(
		"host=%s port=%s user=%s password=%s dbname=%s sslmode=disable TimeZone=America/Sao_Paulo",
		host, port, user, password, dbname,
	)

	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatalf("Erro ao conectar no banco: %v", err)
	}

	sqlDB, _ := db.DB()
	sqlDB.SetMaxOpenConns(10)
	sqlDB.SetMaxIdleConns(5)
	sqlDB.SetConnMaxLifetime(time.Hour)

	DB = db
	log.Println("✅ Banco de dados conectado com sucesso!")

}
