package handlers

import (
	"log"
	"net/http"
	"social_network/internal/http/dto"
	"social_network/internal/service"
	"time"

	"github.com/gin-gonic/gin"
)

type UserHandler struct {
	users service.UserService
}

func NewUserHandler(users service.UserService) *UserHandler {
	return &UserHandler{users: users}
}

func (h *UserHandler) GetByID(c *gin.Context) {
	id := c.Param("id")
	log.Printf("GetByID called: id=%s", id)
	if id == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "missing id"})
		return
	}
	u, err := h.users.GetByID(c.Request.Context(), id)
	if err != nil {
		log.Printf("GetByID error: %v", err)
		c.JSON(http.StatusNotFound, gin.H{"error": "user not found"})
		return
	}
	resp := dto.UserProfileResponse{
		ID:         u.ID,
		FirstName:  u.FirstName,
		SecondName: u.SecondName,
		Birthdate:  u.Birthdate.Format("2006-01-02"),
		Biography:  u.Biography,
		City:       u.City,
	}
	c.JSON(http.StatusOK, resp)
}

func (h *UserHandler) Register(c *gin.Context) {
	var req dto.RegisterUserRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid input"})
		return
	}
	log.Printf("Register called: first_name=%s, second_name=%s", req.FirstName, req.SecondName)
	birth, err := time.Parse("2006-01-02", req.Birthdate)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid birthdate format, expected YYYY-MM-DD"})
		return
	}
	u, err := h.users.CreateProfile(c.Request.Context(), req.FirstName, req.SecondName, birth, req.Biography, req.City, req.Password)
	if err != nil {
		log.Printf("Register error: %v", err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": "internal error"})
		return
	}
	c.JSON(http.StatusOK, dto.RegisterUserResponse{UserID: u.ID})
}
