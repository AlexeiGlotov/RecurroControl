package handler

import (
	"net/http"

	"github.com/gin-gonic/gin"

	"RecurroControl/models"
)

func (h *Handler) getCheat(c *gin.Context) {

	userID, err := getUserId(c)
	if err != nil {
		return
	}

	user, err := h.services.Users.GetUser(userID)

	cheats, err := h.services.Cheats.GetCheats(user.Role)
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err, ErrServerError)
		return
	}

	c.JSON(http.StatusOK, map[string]interface{}{"cheats": cheats})
}

func (h *Handler) createCheat(c *gin.Context) {
	_, err := getUserId(c)
	if err != nil {
		return
	}

	var cheat models.Cheats
	if err := c.BindJSON(&cheat); err != nil {
		newErrorResponse(c, http.StatusBadRequest, err, ErrNotFields)
		return
	}

	id, err := h.services.Cheats.CreateCheat(&cheat)
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err, ErrServerError)
		return
	}

	c.JSON(http.StatusOK, map[string]interface{}{"id": id})
}

func (h *Handler) updateCheat(c *gin.Context) {
	_, err := getUserId(c)
	if err != nil {
		return
	}

	var cheat models.Cheats
	if err := c.BindJSON(&cheat); err != nil {
		newErrorResponse(c, http.StatusBadRequest, err, ErrNotFields)
		return
	}

	err = h.services.Cheats.UpdateCheat(&cheat)
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err, ErrServerError)
		return
	}

	c.Status(http.StatusOK)
}
