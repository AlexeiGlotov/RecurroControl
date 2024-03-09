package handler

import (
	"github.com/gin-gonic/gin"
)

func (h *Handler) createKeys(c *gin.Context) {
	_, err := getUserId(c)
	if err != nil {
		return
	}

	//var inputKeys models.StLicenseKeys

	/*if err := c.BindJSON(&inputKeys); err != nil {
		newErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}*/

	/*cheats, err := h.services.Keys.CreateKeys(inputKeys)
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
	}

	c.JSON(http.StatusOK, map[string]interface{}{"cheats": cheats})*/

}
