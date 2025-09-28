package routes

import (
	"social_network/internal/http/handlers"

	"github.com/gin-gonic/gin"
)

func RegisterUsers(r *gin.RouterGroup, h *handlers.UserHandler) {
	g := r.Group("/user")
	//g.Use(middleware.AuthN()) // требуется токен
	g.GET("get/:id", h.GetByID)
}
