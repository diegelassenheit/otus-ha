package handlers

import (
	"net/http"
	"time"

	"github.com/gin-gonic/gin"

	"social_network/internal/http/dto"
	"social_network/internal/service"
	"social_network/pkg/utils"
)

type AuthHandler struct {
	users service.UserService
}

func NewAuthHandler(users service.UserService) *AuthHandler {
	return &AuthHandler{users: users}
}

func (h *AuthHandler) Register(c *gin.Context) {
	var req dto.RegisterRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid input"})
		return
	}

	// Parse birthdate
	birthdate, err := time.Parse("2006-01-02", req.Birthdate)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid birthdate format, expected YYYY-MM-DD"})
		return
	}

	u, err := h.users.CreateProfile(c.Request.Context(), req.FirstName, req.SecondName, birthdate, req.Biography, req.City, req.Password)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "internal error"})
		return
	}

	c.JSON(http.StatusOK, dto.RegisterUserResponse{UserID: u.ID})
}

func (h *AuthHandler) Login(c *gin.Context) {
	var req dto.LoginRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid input"})
		return
	}

	u, err := h.users.Login(c.Request.Context(), req.ID, req.Password)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "user not found"})
		return
	}

	token, err := utils.GenerateToken(u.ID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "internal error"})
		return
	}

	c.JSON(http.StatusOK, dto.LoginResponse{Token: token})
}
