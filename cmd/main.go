package main

import (
	"go-clean-architecture/api/controllers"
	middlewares "go-clean-architecture/api/middleware"
	"go-clean-architecture/api/routes"
	_ "go-clean-architecture/docs" // Will be generated in step 5
	"go-clean-architecture/internal/auth"
	"go-clean-architecture/internal/ioc"
	"time"

	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"

	"github.com/gin-gonic/gin"
)

// @title Your API Title
// @version 1.0
// @description Your API description

// @securityDefinitions.apikey BearerAuth
// @in header
// @name Authorization
// @description Type "Bearer " followed by a space and your JWT token.
func main() {
	db := ioc.InitializeDB()

	personRepository := ioc.AddPersonRepository(db)
	personService := ioc.AddPersonService(personRepository)
	personController := controllers.NewPersonController(personService)

	jwtCfg := auth.Config{
		Secret:         "CHANGE_ME_USE_ENV",
		AccessTokenTTL: 30 * time.Minute,
		Issuer:         "go-clean-architecture",
	}
	authController := controllers.NewAuthController(jwtCfg)

	router := gin.Default()

	router.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
	routes.RegisterAuthRoutes(router, authController)

	protected := router.Group("")
	protected.Use(middlewares.JWTAuth(jwtCfg))
	routes.RegisterPersonRoutes(protected, personController)

	router.Run(":8000")
}
