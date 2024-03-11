package handler

import (
	"net/http"
	"time"

	"github.com/gin-gonic/gin"

	"RecurroControl/models"
)

// Регистрация

func (h *Handler) signUp(c *gin.Context) {
	var input models.SignUpInput
	if err := c.BindJSON(&input); err != nil {
		newErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}

	if input.Password != input.RePassword {
		newErrorResponse(c, http.StatusBadRequest, "repassword invalid")
		return
	}

	access_key, err := h.services.Authorization.CheckAccessKey(input.Access_Key)
	if err != nil {
		newErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}

	input.Owner = access_key.Owner
	input.Role = access_key.Role

	id, err := h.services.Authorization.CreateUser(input)
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	err = h.services.Authorization.SetLoginAccessKey(input.Login, input.Access_Key)
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
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
		newErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}

	token, err := h.services.Authorization.GenerateToken(input.Login, input.Password)
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
	}

	c.SetCookie("jwt", token, int(time.Hour.Seconds()), "/", "localhost", false, true)

	c.JSON(http.StatusOK, map[string]interface{}{"token": token})

}

func (h *Handler) logout(c *gin.Context) {
	c.SetCookie("jwt", "", -1, "/", "localhost", false, true)
	//c.JSON(http.StatusOK, gin.H{"message": "Logout successful"})

	c.JSON(http.StatusOK, map[string]interface{}{"ok": "ok"})
}

func (h *Handler) renderLoginPage(c *gin.Context) {
	c.HTML(http.StatusOK, "login.html", nil)
}

func (h *Handler) renderRegistrationPage(c *gin.Context) {
	c.HTML(http.StatusOK, "registration.html", nil)
}

func (h *Handler) renderIndexPage(c *gin.Context) {
	c.HTML(http.StatusOK, "index.html", nil)
}
