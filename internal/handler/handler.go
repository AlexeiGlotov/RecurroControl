package handler

import (
	"net/http"

	"github.com/gin-gonic/gin"

	"RecurroControl/internal/service"
)

type Handler struct {
	services *service.Service
}

func NewHandler(services *service.Service) *Handler {
	return &Handler{services: services}
}

func (h *Handler) InitRoutes() *gin.Engine {
	router := gin.New()
	router.Static("/static", "./public")
	router.LoadHTMLGlob("templates/*")

	router.GET("/", func(c *gin.Context) {
		c.HTML(http.StatusOK, "index.html", nil)
	})

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

		cheats := api.Group("/cheats")
		{
			cheats.GET("/", h.getCheat)
		}

		users := api.Group("/users")
		{
			users.GET("/getUserLoginsAndRole", h.getUserLoginsAndRole)
		}
	}

	return router
}
