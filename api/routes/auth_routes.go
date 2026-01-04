package routes

import (
	"go-clean-architecture/api/controllers"

	"github.com/gin-gonic/gin"
)

func RegisterAuthRoutes(r *gin.Engine, authController *controllers.AuthController) {
	r.POST("/auth/login", authController.Login)
}
