package main

import (
	"github.com/gin-gonic/gin"
	"microservicos.com/cmd/config"
	"microservicos.com/cmd/routes"
	entites "microservicos.com/internal/entities"
)

func main() {
	config.Connect()
	config.DB.AutoMigrate(entites.Engresso{}, entites.Administrador{}, entites.DateFormulario{}, entites.Formulario{})
	r := gin.Default()

	routes.Healthy(r)
	routes.CategoryRoutes(r)

	r.Run(":5000")
}
