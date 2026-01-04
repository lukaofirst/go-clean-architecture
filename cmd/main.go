package main

import (
	"go-clean-architecture/api/controllers"
	"go-clean-architecture/api/routes"
	_ "go-clean-architecture/docs" // Will be generated in step 5
	"go-clean-architecture/internal/ioc"

	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"

	"github.com/gin-gonic/gin"
)

// @title Your API Title
// @version 1.0
// @description Your API description
func main() {
	db := ioc.InitializeDB()

	personRepository := ioc.AddPersonRepository(db)
	personService := ioc.AddPersonService(personRepository)
	personController := controllers.NewPersonController(personService)

	router := gin.Default()
	router.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
	routes.RegisterPersonRoutes(router, personController)

	router.Run(":8000")
}
