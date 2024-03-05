package handler

import (
	"github.com/gin-gonic/gin"

	"RecurroControl/pkg/service"
)

type Handler struct {
	services *service.Service
}

func NewHandler(services *service.Service) *Handler {
	return &Handler{services: services}
}

func (h *Handler) InitRoutes() *gin.Engine {
	router := gin.New()

	auth := router.Group("/auth")
	{
		auth.POST("/sign-up", h.signUp)
		auth.POST("/sign-in", h.signIn)
		auth.POST("/check-admission-key", h.checkAdmissionKey)
	}

	api := router.Group("/api", h.userIdentity)
	{

		admission := api.Group("/admission")
		{
			admission.POST("/", h.createKey)
			admission.GET("/", h.getKey)

		}

	}

	return router
}
