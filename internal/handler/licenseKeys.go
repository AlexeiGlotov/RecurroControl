package handler

import (
	"net/http"

	"github.com/gin-gonic/gin"

	"RecurroControl/models"
)

func (h *Handler) createLicenseKeys(c *gin.Context) {
	_, err := getUserId(c)
	if err != nil {
		return
	}

	var cheat []models.LicenseKeys
	if err := c.BindJSON(&cheat); err != nil {
		newErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}

	/*id, err := h.services.Cheats.CreateCheats(&cheat)
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
	}

	c.JSON(http.StatusOK, map[string]interface{}{"id": id})*/
}
