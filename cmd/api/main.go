package main

import (
	"log"
	"os"

	"github.com/gin-gonic/gin"
	handler "github.com/matrcross/digital-wallet-api/internal/http/user"
	"github.com/matrcross/digital-wallet-api/internal/user"
	"github.com/matrcross/digital-wallet-api/internal/user/mock"
)

func main() {
	router := gin.New()

	repo := mock.NewMockRepository()
	service := user.NewService(repo)
	handler := handler.NewHandler(service)
	handler.RegisterRoutes(router)

	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}
	if err := router.Run(":" + port); err != nil {
		log.Fatalf("Failed to start server: %v", err)
	}
}
