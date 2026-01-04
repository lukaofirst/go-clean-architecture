package controllers

import (
	"net/http"
	"time"

	"go-clean-architecture/internal/auth"

	"github.com/gin-gonic/gin"
)

type AuthController struct {
	cfg auth.Config
}

func NewAuthController(cfg auth.Config) *AuthController {
	return &AuthController{cfg: cfg}
}

type LoginRequest struct {
	Username string `json:"username" binding:"required"`
	Password string `json:"password" binding:"required"`
}

type LoginResponse struct {
	AccessToken string `json:"accessToken"`
	ExpiresIn   int64  `json:"expiresInSeconds"`
	TokenType   string `json:"tokenType"`
}

// Login godoc
// @Summary Login and get JWT
// @Description Authenticates a user and returns an access token
// @Tags auth
// @Accept json
// @Produce json
// @Param body body LoginRequest true "Credentials"
// @Success 200 {object} LoginResponse
// @Failure 400 {object} map[string]any
// @Failure 401 {object} map[string]any
// @Router /auth/login [post]
func (a *AuthController) Login(c *gin.Context) {
	var req LoginRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// TODO: Replace with your real user validation (DB/service)
	if req.Username != "admin" || req.Password != "admin" {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "invalid credentials"})
		return
	}

	token, err := auth.GenerateAccessToken(a.cfg, "user-123")
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "could not generate token"})
		return
	}

	c.JSON(http.StatusOK, LoginResponse{
		AccessToken: token,
		ExpiresIn:   int64((30 * time.Minute).Seconds()),
		TokenType:   "Bearer",
	})
}
