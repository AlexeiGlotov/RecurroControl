package handler

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func (h *Handler) getUsers(c *gin.Context) {
	_, err := getUserId(c)
	if err != nil {
		return
	}

	cheats, err := h.services.Users.GetUsers()
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
	}

	c.JSON(http.StatusOK, map[string]interface{}{"Users": cheats})

}
