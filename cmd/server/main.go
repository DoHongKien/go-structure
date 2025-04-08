package main

import (
	"github.com/DoHongKien/go-structure/internal/initialize"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

// @title Example API
// @version 1.0
// @description This is a sample server for demonstrating Swagger with Go.
// @host localhost:9999
// @BasePath /

// @Summary Get Swagger documentation
// @Description Serve Swagger documentation for the API.
// @Tags Swagger
// @Accept json
// @Produce json
// @Success 200 {object} map[string]interface{}
// @Router /swagger/*any [get]
func main() {

	r := initialize.Run()

	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	r.Run(":9999")
}
