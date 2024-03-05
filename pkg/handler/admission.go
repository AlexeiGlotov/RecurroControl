package handler

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func (h *Handler) createKey(c *gin.Context) {
	userID, err := getUserId(c)
	if err != nil {
		return
	}

	key, err := h.services.Admission.CreateKey(userID)
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}
	c.JSON(http.StatusOK, map[string]interface{}{
		"key": key,
	})

}

func (h *Handler) getKey(c *gin.Context) {
	_, err := getUserId(c)
	if err != nil {
		return
	}

	key, err := h.services.Admission.GetKey()
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}
	c.JSON(http.StatusOK, map[string]interface{}{
		"key": key,
	})
}
