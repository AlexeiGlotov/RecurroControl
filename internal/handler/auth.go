package handler

import (
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"

	"RecurroControl/models"
)

// Регистрация

func (h *Handler) signUp(c *gin.Context) {
	var input models.SignUpInput

	if err := c.BindJSON(&input); err != nil {
		newErrorResponse(c, http.StatusBadRequest, err, ErrNotFields)
		return
	}

	if len(input.Login) > 20 || len(input.Login) < 4 {
		newErrorResponse(c, http.StatusBadRequest, nil, "bad len login")
		return
	}

	if len(input.Password) > 20 || len(input.Password) < 6 {
		newErrorResponse(c, http.StatusBadRequest, nil, "bad len password")
		return
	}

	if input.Password != input.RePassword {
		newErrorResponse(c, http.StatusBadRequest, nil, "check the repassword != passwords are correct")
		return
	}

	access_key, err := h.services.Authorization.CheckAccessKey(input.Access_Key)
	if err != nil {
		newErrorResponse(c, http.StatusBadRequest, err, "invalid key")
		return
	}

	input.Owner = access_key.Owner
	input.Role = access_key.Role

	id, err := h.services.Authorization.CreateUser(input)
	if err != nil {

		if strings.Contains(err.Error(), "Duplicate entry") {
			newErrorResponse(c, http.StatusBadRequest, err, "enter another login")
		} else {
			newErrorResponse(c, http.StatusInternalServerError, err, "create user error")
		}

		return
	}

	err = h.services.Authorization.SetLoginAccessKey(input.Login, input.Access_Key)
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err, "error at the end of registration")
		return
	}

	c.JSON(http.StatusOK, map[string]interface{}{"id": id})
}

type signInInput struct {
	Login    string `json:"username" binding:"required"`
	Password string `json:"password" binding:"required"`
}

// Авторизация
func (h *Handler) signIn(c *gin.Context) {

	var input signInInput

	if err := c.BindJSON(&input); err != nil {
		newErrorResponse(c, http.StatusBadRequest, err, ErrNotFields)
		return
	}

	token, err := h.services.Authorization.GenerateToken(input.Login, input.Password)
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err, ErrServerError)
		return
	}

	c.JSON(http.StatusOK, map[string]interface{}{"token": token})
}
