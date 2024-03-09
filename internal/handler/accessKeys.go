package handler

import (
	"net/http"

	"github.com/gin-gonic/gin"

	"RecurroControl/models"
)

func (h *Handler) createAccessKeys(c *gin.Context) {
	userID, err := getUserId(c)
	if err != nil {
		return
	}

	var access_key models.AccessKey
	if err := c.BindJSON(&access_key); err != nil {
		newErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}

	key, err := h.services.AccessKeys.CreateAccessKey(userID, access_key.Role)
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}
	c.JSON(http.StatusOK, map[string]interface{}{
		"key": key,
	})

}

func (h *Handler) getAccessKey(c *gin.Context) {
	userID, err := getUserId(c)
	if err != nil {
		return
	}

	user, err := h.services.Users.GetUserStruct(userID)
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	key, err := h.services.AccessKeys.GetAccessKey(user.Login, user.Role)
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}
	c.JSON(http.StatusOK, map[string]interface{}{
		"key": key,
	})
}
