// Package main implements a task management API.
//
//	Schemes: http
//	BasePath: /
//	Version: 1.0.0
//
//	Consumes:
//	- application/json
//
//	Produces:
//	- application/json
//
// swagger:meta
package main

import (
	"log"
	"os"
	"prova-fattocs/internal/config"
	"prova-fattocs/internal/infra/database"
	"prova-fattocs/internal/routes"

	cors "github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

// @title Task Management API
// @version 1.0
// @description This is a simple task management API.
// @contact.name API Support
// @contact.email support@example.com
// @host localhost:8080
// @BasePath /
func main() {
	if mode := os.Getenv("GIN_MODE"); mode == "release" {
		gin.SetMode(gin.ReleaseMode)
	} else {
		gin.SetMode(gin.DebugMode)
	}

	cfg := config.Load()
	db := database.NewPostgresConnection(cfg)
	r := gin.Default()

	corsConfig := cors.DefaultConfig()
	corsConfig.AllowAllOrigins = true
	corsConfig.AllowMethods = []string{"GET", "POST", "PUT", "PATCH", "DELETE", "HEAD"}
	corsConfig.AllowHeaders = []string{"Origin", "Content-Length", "Content-Type", "Authorization"}
	r.Use(cors.New(corsConfig))

	r.ForwardedByClientIP = false

	// @Summary Show API documentation
	// @Description Show API documentation with Swagger UI
	// @Tags Documentation
	// @Accept */*
	// @Produce html
	// @Success 200
	// @Router /swagger/*any [get]
	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	routes.SetupRoutes(r, db)

	log.Fatal(r.Run(":" + cfg.ServerPort))
}
