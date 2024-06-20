package controller

import (
	"fmt"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
	"microservicos.com/cmd/config"
	entites "microservicos.com/internal/entities"
	//"microservicos.com/pkg/util"
)

// Login handles user login and returns a JWT token if credentials are correct
func Login(ctx *gin.Context) {
	var p entites.Login
	var user entites.Administrador

	db := config.GetDatabase()

	if err := ctx.ShouldBindJSON(&p); err != nil {
		ctx.JSON(400, gin.H{
			"error": "cannot bind JSON: " + err.Error(),
		})
		return
	}
	// tentei transforma em um middlewares mas não conseguir, tentar novamente pelo turno da noite
	// Check if the user with the provided CPF exists in the database
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
	fmt.Print(user.Password)

	ctx.JSON(200, gin.H{
		"token": user,
	})
}
