package controller

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
	"microservicos.com/cmd/config"
	entites "microservicos.com/internal/entities"
	//"microservicos.com/pkg/util"
)

func LoginAdm(ctx *gin.Context) {
	var p entites.Login
	var user entites.Administrador

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

	ctx.JSON(http.StatusOK, gin.H{
		"mensagem": "Login realizado com sucesso",
	})
}
