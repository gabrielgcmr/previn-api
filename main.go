package main

import (
	"log"

	"github.com/gabrielgcmr/previn-api/internal/database"
	"github.com/gabrielgcmr/previn-api/internal/middleware"
	"github.com/gabrielgcmr/previn-api/internal/patient"
	"github.com/gabrielgcmr/previn-api/pkg/validation"

	"github.com/gin-gonic/gin"
)

func main() {
	//conectar db
	database.Connect()
	database.DB.AutoMigrate(&patient.Patient{})
	//montar o gin e rotas
	r := gin.Default()

	// 🌐 Aplica o middleware de CORS
	r.Use(middleware.SetupCors())

	// Registra as rotas
	patient.Routes(r)

	_ = validation.Init()

	if err := validation.Init(); err != nil {
		log.Fatalf("Erro ao iniciar validador: %v", err)
	}

	log.Println("🚀 API running at http://localhost:8080")
	r.Run(":8080")
}
