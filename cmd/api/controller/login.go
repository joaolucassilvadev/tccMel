package controller

import (
	"fmt"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
	"microservicos.com/cmd/config"
	entites "microservicos.com/internal/entities"
	//"microservicos.com/pkg/util"
)

// Login godoc
// @Summary Login do usuário
// @Description Realiza o login do engresso com cpf e data.
// @Tags auth
// @Accept  json
// @Produce  json
// @Param   login  body  entites.Loginengresso  true  "Dados de Login"
// @Router /login [post]
func Login(ctx *gin.Context) {
	now := time.Now()

	var p entites.Loginengresso
	var user entites.Engresso
	var dataForm entites.DateFormulario

	db := config.GetDatabase()

	if err := ctx.ShouldBindJSON(&p); err != nil {
		ctx.JSON(400, entites.ErrorResponse{
			Error: "cannot bind JSON: " + err.Error(),
		})
		return
	}

	if err := db.Where("cpf= ?", p.Cpf).First(&user).Error; err != nil {
		fmt.Printf("nome do email %s", p.Cpf)
		if err == gorm.ErrRecordNotFound {
			ctx.JSON(400, entites.ErrorResponse{
				Error: "user not found",
			})
		} else {
			ctx.JSON(500, entites.ErrorResponse{
				Error: "database error: " + err.Error(),
			})
		}
		return
	}
	config.DB.Find(&dataForm)
	fmt.Print(dataForm.DateFim)

	fim, err := time.Parse("2006-01-02", dataForm.DateFim)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, entites.ErrorResponse{Error: "Invalid date_fim format"})
		return
	}

	if fim.Before(now) {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"mensagem de erro": "A data para inscrição do formulário já passou",
		})
		return
	}

	ctx.JSON(http.StatusOK, entites.LoginSuccessResponse{
		Mensagem: "Login realizado com sucesso",
	})
}
