package controller

import (
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"microservicos.com/cmd/config"
	entites "microservicos.com/internal/entities"
)

type date struct {
	datainicio time.Time
	datafim    time.Time
}

func DataController(ctx *gin.Context) {
	var formDTO entites.DateFormulario

	// Bind JSON request body to DTO
	if err := ctx.ShouldBindJSON(&formDTO); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Por favor, forneça dados válidos"})
		return
	}

	// Converter DTO para entidade
	formEntity := entites.DateFormulario{
		DateInicio: formDTO.DateInicio,
		DateFim:    formDTO.DateFim,
	}

	// Salvar no banco de dados usando GORM
	if err := config.DB.Create(&formEntity).Error; err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": "Falha ao salvar no banco de dados", "details": err.Error()})
		return
	}

	// Responder com sucesso
	ctx.JSON(http.StatusOK, gin.H{
		"message":     "Formulário criado com sucesso",
		"data_inicio": formEntity.DateInicio,
		"data_fim":    formEntity.DateFim,
	})
}
