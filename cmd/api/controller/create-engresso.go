package controller

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"microservicos.com/internal/repositories"
	usecases "microservicos.com/internal/use-cases"
)

type createCategoryInput struct {
	Name            string `json:"name" binding:"required"`
	Cpf             string `json:"cpf" binding:"required"`
	Data_Nascimento string `json:"data_nascimento"`
}

// Login godoc
// @Summary create engresso
// @Description para realizar o cadastros do engresso precisa name,cpf,data nascimento.
// @Tags auth
// @Accept  json
// @Produce  json
// @Param   login  body  entites.Engresso  true  "Dados ebresso"
// @Router /categories/createCategory [post]
func CreateCategory(ctx *gin.Context, repository repositories.IRepositories) {
	var body createCategoryInput
	if err := ctx.ShouldBindJSON(&body); err != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"message error": err.Error(),
		})
		return
	}
	useCase := usecases.NewcreateCategoryUseCase(repository)
	err := useCase.Execute(body.Name, body.Cpf, body.Data_Nascimento)
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
