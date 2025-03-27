package router

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func NewRouter() *gin.Engine {

	r := gin.Default()

	v2 := r.Group("/api/v1")
	{
		v2.GET("/users", PONG)
		v2.GET("/users/:id", PONG)
	}

	return r
}

func PONG(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"message": "pong",
	})
}
