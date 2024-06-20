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

func Login(ctx *gin.Context) {
	now := time.Now()

	var p entites.Login
	var user entites.Administrador
	var dataForm entites.DateFormulario

	db := config.GetDatabase()

	if err := ctx.ShouldBindJSON(&p); err != nil {
		ctx.JSON(400, gin.H{
			"error": "cannot bind JSON: " + err.Error(),
		})
		return
	}

	if err := db.Where("email = ?", p.Email).First(&user).Error; err != nil {
		fmt.Printf("nome do email %s", p.Email)
		if err == gorm.ErrRecordNotFound {
			ctx.JSON(400, gin.H{
				"error": "user not found",
			})
		} else {
			ctx.JSON(500, gin.H{
				"error": "database error: " + err.Error(),
			})
		}
		return
	}
	config.DB.Find(&dataForm)
	fmt.Print(dataForm.DateFim)

	fim, err := time.Parse("2006-01-02", dataForm.DateFim)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Invalid date_fim format"})
		return
	}

	if fim.Before(now) {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"mensagem de erro": "A data para inscrição do formulário já passou",
		})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{
		"mensagem": "Login realizado com sucesso",
	})
}
