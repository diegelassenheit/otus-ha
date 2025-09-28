package http

import (
	"social_network/internal/http/handlers"
	"social_network/internal/http/routes"

	"github.com/gin-gonic/gin"
)

func NewRouter(auth *handlers.AuthHandler, users *handlers.UserHandler) *gin.Engine {
	r := gin.New()
	r.Use(gin.Recovery())

	r.GET("/healthz", func(c *gin.Context) { c.Status(204) })

	v1 := r.Group("/")
	routes.RegisterAuth(v1, auth)
	routes.RegisterUsers(v1, users)
	//routes.RegisterFeed(v1, d.Feed)

	return r
}
