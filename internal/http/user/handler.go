package handler

import (
	"github.com/gin-gonic/gin"
	"github.com/matrcross/digital-wallet-api/internal/user"
)

type Handler struct {
	userService *user.Service
}

func NewHandler(userService *user.Service) *Handler {
	return &Handler{userService: userService}
}

func (h *Handler) RegisterRoutes(router *gin.Engine) {
	router.GET("/user/:id", h.Get)
	router.POST("/user/", h.Create)
	router.POST("/user/:id/deposit", h.Deposit)
}
