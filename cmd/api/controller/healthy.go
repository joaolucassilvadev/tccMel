package controller

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

// Login godoc
// @Summary service on
// @Description essa rota server pra saber se nosso service tá on.
// @Tags auth
// @Accept  json
// @Produce  json
// @Router /healthy [get]
func Healthy(ctx *gin.Context) {
	ctx.JSON(http.StatusOK, gin.H{
		"mensage servidor on": "Bomm dia mellllllllllllllllllllllllllllllllllllllllllllll",
	})
}
