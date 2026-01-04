package controllers

import (
	"go-clean-architecture/internal/application/services"
	"go-clean-architecture/internal/domain/entities"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

type PersonController interface {
	GetAll(ctx *gin.Context)
	GetByID(ctx *gin.Context)
	Create(ctx *gin.Context)
	Update(ctx *gin.Context)
	Delete(ctx *gin.Context)
}

type personController struct {
	PersonService services.PersonService
}

func NewPersonController(personService services.PersonService) PersonController {
	return &personController{PersonService: personService}
}

// GetAll godoc
// @Security BearerAuth
// @Summary Get all persons
// @Description Get a list of all persons in the system
// @Tags persons
// @Produce json
// @Success 200 {array} entities.Person
// @Router /persons [get]
func (c *personController) GetAll(ctx *gin.Context) {
	persons, err := c.PersonService.GetAll()
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	ctx.JSON(http.StatusOK, persons)
}

// GetByID godoc
// @Security BearerAuth
// @Summary Get a person by ID
// @Description Get details of a single person by their ID
// @Tags persons
// @Produce  json
// @Param id path int true "Person ID"
// @Success 200 {object} entities.Person
// @Failure 404 {object} map[string]string
// @Router /persons/{id} [get]
func (c *personController) GetByID(ctx *gin.Context) {
	id, err := strconv.Atoi(ctx.Param("id"))
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Invalid ID"})
		return
	}

	person, err := c.PersonService.GetByID(uint(id))
	if err != nil {
		ctx.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, person)
}

func (c *personController) Create(ctx *gin.Context) {
	var person entities.Person
	if err := ctx.ShouldBindJSON(&person); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if err := c.PersonService.Create(person); err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusCreated, person)
}

func (c *personController) Update(ctx *gin.Context) {
	var person entities.Person
	if err := ctx.ShouldBindJSON(&person); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if err := c.PersonService.Update(person); err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, person)
}

func (c *personController) Delete(ctx *gin.Context) {
	id, err := strconv.Atoi(ctx.Param("id"))
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Invalid ID"})
		return
	}

	if err := c.PersonService.Delete(uint(id)); err != nil {
		ctx.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
		return
	}

	ctx.Status(http.StatusNoContent)
}
