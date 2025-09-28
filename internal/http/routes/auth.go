package routes

import (
	"social_network/internal/http/handlers"

	"github.com/gin-gonic/gin"
)

func RegisterAuth(r *gin.RouterGroup, h *handlers.AuthHandler) {
	// Login at root level to match OpenAPI spec
	r.POST("/login", h.Login)

	// Register under /user group
	userGroup := r.Group("/user")
	userGroup.POST("/register", h.Register)
}
