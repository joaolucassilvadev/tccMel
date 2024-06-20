package routes

import (
	"github.com/gin-gonic/gin"
	"microservicos.com/cmd/api/controller"
	"microservicos.com/internal/repositories"
)

func CategoryRoutes(r *gin.Engine) {
	categoryRoutes := r.Group("/categories")
	categoryRepository := repositories.NewCategoryRepository()

	categoryRoutes.POST("/createCategory", func(ctx *gin.Context) {
		controller.CreateCategory(ctx, categoryRepository)
	})

	adminRoutes := r.Group("/administrators")
	administradorRepository := repositories.NewAdministradorRepository()

	adminRoutes.POST("/createAdministrator", func(ctx *gin.Context) {
		controller.CreateAdmistrador(ctx, administradorRepository)
	})
	r.POST("/login", controller.Login)
	r.POST("/dateformulario", controller.DataController)
	r.POST("/formulario", controller.CreateFormularioController)
}

func Healthy(r *gin.Engine) {
	r.GET("/healthy", controller.Healthy)
}
