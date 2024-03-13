package handler

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func (h *Handler) getUserLoginsAndRole(c *gin.Context) {
	userID, err := getUserId(c)
	if err != nil {
		return
	}

	cheats, err := h.services.Users.GetUserLoginsAndRole(userID)
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
	}

	c.JSON(http.StatusOK, map[string]interface{}{"users": cheats})
}
