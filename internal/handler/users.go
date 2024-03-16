package handler

import (
	"net/http"

	"github.com/gin-gonic/gin"

	"RecurroControl/models"
)

func (h *Handler) getUsers(c *gin.Context) {
	userID, err := getUserId(c)
	if err != nil {
		return
	}

	user, err := h.services.Users.GetUser(userID)
	if user.Role == models.Salesman {
		newErrorResponse(c, http.StatusForbidden, err, ErrAccessDenied)
		return
	}

	users, err := h.services.Users.GetUsers(userID)
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err, ErrServerError)
		return
	}

	c.JSON(http.StatusOK, map[string]interface{}{"users": users})
}

type inputUserAction struct {
	Id int `json:"id" binding:"required"`
}

func (h *Handler) getUser(c *gin.Context) {
	userID, err := getUserId(c)
	if err != nil {
		return
	}

	var input inputUserAction
	if err := c.BindJSON(&input); err != nil {
		newErrorResponse(c, http.StatusBadRequest, err, ErrNotFields)
		return
	}

	user, err := h.services.Users.GetUser(userID)
	if user.Role == models.Salesman {
		newErrorResponse(c, http.StatusForbidden, err, ErrAccessDenied)
		return
	}

	cheats, err := h.services.Users.GetUser(input.Id)
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err, ErrServerError)
		return
	}

	c.JSON(http.StatusOK, map[string]interface{}{"user": cheats})
}

func (h *Handler) ban(c *gin.Context) {
	userID, err := getUserId(c)
	if err != nil {
		return
	}

	var input inputUserAction
	if err := c.BindJSON(&input); err != nil {
		newErrorResponse(c, http.StatusBadRequest, err, ErrNotFields)
		return
	}

	user, err := h.services.Users.GetUser(userID)
	if user.Role == models.Salesman {
		newErrorResponse(c, http.StatusForbidden, err, ErrAccessDenied)
		return
	}

	err = h.services.Users.Ban(input.Id)
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err, ErrServerError)
		return
	}

	c.Status(http.StatusOK)
}

func (h *Handler) unban(c *gin.Context) {
	userID, err := getUserId(c)
	if err != nil {
		return
	}

	var input inputUserAction
	if err := c.BindJSON(&input); err != nil {
		newErrorResponse(c, http.StatusBadRequest, err, ErrNotFields)
		return
	}

	user, err := h.services.Users.GetUser(userID)
	if user.Role == models.Salesman {
		newErrorResponse(c, http.StatusForbidden, err, ErrAccessDenied)
		return
	}

	err = h.services.Users.Unban(input.Id)
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err, ErrServerError)
		return
	}

	c.Status(http.StatusOK)
}

func (h *Handler) delete(c *gin.Context) {
	userID, err := getUserId(c)
	if err != nil {
		return
	}

	var input inputUserAction
	if err := c.BindJSON(&input); err != nil {
		newErrorResponse(c, http.StatusBadRequest, err, ErrNotFields)
		return
	}

	user, err := h.services.Users.GetUser(userID)
	if user.Role != models.Admin {
		newErrorResponse(c, http.StatusForbidden, err, ErrAccessDenied)
		return
	}

	err = h.services.Users.Delete(input.Id)
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err, ErrServerError)
		return
	}

	c.Status(http.StatusOK)
}
