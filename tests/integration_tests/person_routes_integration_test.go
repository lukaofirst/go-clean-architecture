package integration_tests

import (
	"encoding/json"
	"go-clean-architecture/api/controllers"
	"go-clean-architecture/api/routes"
	"go-clean-architecture/internal/domain/entities"
	infrastructure "go-clean-architecture/internal/infrastructure/persistence"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
)

func TestPersonRoutes(t *testing.T) {
	db, teardown, err := setupTestContainer()
	if err != nil {
		t.Fatalf("Failed to set up test container: %v", err)
	}
	defer teardown()

	// Set up Gin router and dependencies
	gin.SetMode(gin.TestMode)
	router := gin.Default()

	repo := infrastructure.NewPersonRepository(db) // Use the GORM DB
	controller := controllers.NewPersonController(repo)
	routes.RegisterPersonRoutes(router, controller)

	// Integration test: GET /persons
	t.Run("GetAllPersons", func(t *testing.T) {
		req, _ := http.NewRequest(http.MethodGet, "/persons", nil)
		resp := httptest.NewRecorder()

		router.ServeHTTP(resp, req)

		assert.Equal(t, http.StatusOK, resp.Code)
		assert.Contains(t, resp.Body.String(), "[]") // Should return empty array initially
	})

	// Integration test: POST /persons
	t.Run("CreatePerson", func(t *testing.T) {
		person := entities.Person{
			ID:    1,
			Name:  "John Doe",
			Age:   30,
			Email: "lorem@gmail.com",
		}
		body, _ := json.Marshal(person)
		req, _ := http.NewRequest(http.MethodPost, "/persons", strings.NewReader(string(body)))
		req.Header.Set("Content-Type", "application/json")
		resp := httptest.NewRecorder()

		router.ServeHTTP(resp, req)

		assert.Equal(t, http.StatusCreated, resp.Code)
	})
}
