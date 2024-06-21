package main

import (
	"log"
	"os"

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

	host := os.Getenv("HOST")
	if host == "" {
		host = "0.0.0.0"
	}

	port := os.Getenv("PORT")
	if port == "" {
		port = "8089"
	}

	address := host + ":" + port
	if err := r.Run(address); err != nil {
		log.Panicf("error: %s", err)
	}
}
