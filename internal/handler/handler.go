package handler

import (
	"github.com/gin-contrib/cors"
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
	router := gin.Default()

	router.Use(cors.New(cors.Config{
		//AllowOrigins:     []string{"http://localhost:3000"},
		AllowMethods:     []string{"GET", "POST", "PUT", "PATCH", "DELETE", "OPTIONS"},
		AllowHeaders:     []string{"Origin", "Content-Type", "Authorization"},
		AllowCredentials: true,
		AllowAllOrigins:  true,
	}))

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
			cheats.PUT("/", h.updateCheat)
		}

		users := api.Group("/users")
		{
			users.GET("/getUsers", h.getUsers)
			users.POST("/getUser", h.getUser)
			users.POST("/ban", h.ban)
			users.POST("/unban", h.unban)
			users.POST("/delete", h.delete)
		}

		licenseKeys := api.Group("/license-keys")
		{
			licenseKeys.POST("/", h.createLicenseKeys)
			licenseKeys.GET("/", h.getLicenseKeys)
			licenseKeys.GET("/getCustomLicenseKey", h.getCustomLicenseKeys)
			licenseKeys.POST("/resetHWID", h.licenseKeyResetHWID)
			licenseKeys.POST("/ban", h.licenseKeyBan)
			licenseKeys.POST("/ban-of-date", h.licenseKeysBanDate)
			licenseKeys.POST("/unban", h.licenseKeyUnban)
			licenseKeys.POST("/delete", h.licenseKeyDelete)
		}
	}

	return router
}
