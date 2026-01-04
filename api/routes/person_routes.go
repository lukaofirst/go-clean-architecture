package routes

import (
	api "go-clean-architecture/api/controllers"

	"github.com/gin-gonic/gin"
)

func RegisterPersonRoutes(r gin.IRoutes, personController api.PersonController) {
	r.GET("/persons", personController.GetAll)
	r.GET("/persons/:id", personController.GetByID)
	r.POST("/persons", personController.Create)
	r.PUT("/persons/:id", personController.Update)
	r.DELETE("/persons/:id", personController.Delete)
}
