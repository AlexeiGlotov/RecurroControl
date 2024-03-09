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
	}

	api := router.Group("/api", h.userIdentity)
	{
		accessKeys := api.Group("/access-keys")
		{
			accessKeys.POST("/", h.createAccessKeys)
			accessKeys.GET("/", h.getAccessKey)
		}

		cheats := api.Group("/cheats")
		{
			cheats.POST("/", h.createCheat)
			cheats.GET("/", h.getCheat)
		}

		users := api.Group("/users")
		{
			users.GET("/getUserLoginsAndRole", h.getUserLoginsAndRole)
		}

		licenseKeys := api.Group("/license-keys")
		{
			licenseKeys.POST("/", h.createLicenseKeys)
		}
	}

	return router
}
