package handler

import (
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

	logic := router.Group("/", h.authMiddleware())
	{
		logic.GET("/", h.renderIndexPage)
	}

	router.GET("/login", h.noAuthMiddleware(), h.renderLoginPage)
	router.GET("/registration", h.noAuthMiddleware(), h.renderRegistrationPage)

	router.GET("/logout", h.logout)
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
