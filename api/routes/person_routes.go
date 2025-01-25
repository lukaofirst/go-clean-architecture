package routes

import (
	api "go-clean-architecture/api/controllers"

	"github.com/gin-gonic/gin"
)

func RegisterPersonRoutes(router *gin.Engine, personController api.PersonController) {
	router.GET("/persons", personController.GetAll)
	router.GET("/persons/:id", personController.GetByID)
	router.POST("/persons", personController.Create)
	router.PUT("/persons/:id", personController.Update)
	router.DELETE("/persons/:id", personController.Delete)
}
