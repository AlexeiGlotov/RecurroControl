package handler

import (
	"net/http"

	"github.com/gin-gonic/gin"

	todo "RecurroControl"
)

// Регистрация

func (h *Handler) signUp(c *gin.Context) {
	var input todo.SignUpInput
	if err := c.BindJSON(&input); err != nil {
		newErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}

	if input.Password != input.RePassword {
		newErrorResponse(c, http.StatusBadRequest, "repassword invalid")
		return
	}

	// проверить что в таблице ключей есть такой ключ и вернуть owner
	owner, err := h.services.CheckKeyAdmission(input.Access_Key)
	if err != nil {
		newErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}

	input.Owner = owner

	id, err := h.services.Authorization.CreateUser(input)
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	// Запись новое значение в isLogin в таблицу ключей

	c.JSON(http.StatusOK, map[string]interface{}{"id": id})
}

type signInInput struct {
	Login    string `json:"login" binding:"required"`
	Password string `json:"password" binding:"required"`
}

// Авторизация
func (h *Handler) signIn(c *gin.Context) {
	var input signInInput

	if err := c.BindJSON(&input); err != nil {
		newErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}

	token, err := h.services.Authorization.GenerateToken(input.Login, input.Password)
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
	}

	c.JSON(http.StatusOK, map[string]interface{}{"token": token})

}

type stKey struct {
	Key string `json:"key"`
}

func (h *Handler) checkAdmissionKey(c *gin.Context) {

	key := stKey{}
	if err := c.BindJSON(&key); err != nil {
		newErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}

	owner, err := h.services.CheckKeyAdmission(key.Key)
	if err != nil {
		newErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}
	c.JSON(http.StatusOK, map[string]interface{}{
		"owner": owner,
	})
}
