package handler

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func (h *Handler) getCheat(c *gin.Context) {
	_, err := getUserId(c)
	if err != nil {
		return
	}

	cheats, err := h.services.Cheat.GetCheats()
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
	}

	c.JSON(http.StatusOK, map[string]interface{}{"cheats": cheats})

}
