package main

import (
	"go-clean-architecture/api/controllers"
	"go-clean-architecture/api/routes"
	"go-clean-architecture/internal/ioc"

	"github.com/gin-gonic/gin"
)

func main() {
	db := ioc.InitializeDB()

	personRepository := ioc.AddPersonRepository(db)
	personService := ioc.AddPersonService(personRepository)
	personController := controllers.NewPersonController(personService)

	router := gin.Default()
	routes.RegisterPersonRoutes(router, personController)

	router.Run(":8000")
}
