package controller

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"microservicos.com/cmd/config"
	entities "microservicos.com/internal/entities"
)

// DataController godoc
// @Summary Editar a data de fim de um formulário
// @Description Permite editar a data de fim de um formulário existente.
// @Tags formulario
// @Accept json
// @Produce json
// @Param data body entites.DateFormulario true "Dados para edição da data de fim do formulário"
// @Router /EditData [put]
func EditData(ctx *gin.Context) {
	var formDTO entities.DateFormulario

	if err := ctx.ShouldBindJSON(&formDTO); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Por favor, forneça dados válidos"})
		return
	}

	var existingForm entities.DateFormulario
	if err := config.DB.First(&existingForm).Error; err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Formulário não encontrado"})
		return
	}

	existingForm.DateInicio = formDTO.DateInicio
	existingForm.DateFim = formDTO.DateFim

	if err := config.DB.Save(&existingForm).Error; err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": "Falha ao salvar no banco de dados", "details": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{
		"message":     "Formulário atualizado com sucesso",
		"data_inicio": existingForm.DateInicio,
		"data_fim":    existingForm.DateFim,
	})
}
