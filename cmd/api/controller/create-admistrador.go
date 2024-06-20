package controller

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"microservicos.com/internal/repositories"
	usecases "microservicos.com/internal/use-cases"
)

// aqui vamos utilizar como dto, por isso está em minusculo
type createLogin struct {
	Name string `json:"name" binding:"required"`

	Email string `json:"Email" binding:"required"`
	senha string `json:"senha"`
}

func CreateAdmistrador(ctx *gin.Context, repository repositories.IRepositoriess) {
	var body createLogin
	if err := ctx.ShouldBindJSON(&body); err != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"message error": err.Error(),
		})
		return
	}
	useCase := usecases.NewcreateAdmistradorUseCase(repository)
	err := useCase.Execute(body.Name, body.Email, body.senha)
	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"success": false,
			"error":   err.Error(),
		})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{
		"success": true,
		"message": "Cadastrado com sucesso, o email cadastrado no portal será o email a qual será enviado o contracheque e outros recebibos.",
	})
}
