package main

import (
	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"
	"github.com/swaggo/gin-swagger"
	"microservicos.com/cmd/config"
	"microservicos.com/cmd/routes"
	_ "microservicos.com/docs" // Importa os documentos do Swagger
	entites "microservicos.com/internal/entities"
)

// @title Microservices API
// @version 1.0
// @description Esta é a documentação da API dos microserviços.
// @termsOfService http://swagger.io/terms/

// @contact.name API Support
// @contact.url http://www.swagger.io/support
// @contact.email support@swagger.io

// @license.name Apache 2.0
// @license.url http://www.apache.org/licenses/LICENSE-2.0.html

// @BasePath /

// @securityDefinitions.basic BasicAuth
func main() {
	config.Connect()
	config.DB.AutoMigrate(entites.Engresso{}, entites.Administrador{}, entites.DateFormulario{}, entites.Formulario{})
	r := gin.Default()

	// Configurações do Swagger
	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	routes.Healthy(r)
	routes.CategoryRoutes(r)

	r.Run(":5000")
}
