package handler

import (
	"errors"
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
)

const (
	authorizationHeader = "Authorization"
	userCtx             = "userID"
)

func (h *Handler) ensureNotLoggedIn() gin.HandlerFunc {
	return func(c *gin.Context) {
		if _, err := c.Cookie("jwt"); err == nil {
			c.Redirect(http.StatusSeeOther, "/")
			c.Abort()
			return
		}
		c.Next()
	}
}

func (h *Handler) authMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		tokenString, err := c.Cookie("jwt")
		if err != nil {
			//c.JSON(http.StatusMovedPermanently, gin.H{"error": "Unauthorized"})
			c.Redirect(301, "/login")
			c.Abort()
			return
		}

		_, err = h.services.Authorization.ParseToken(tokenString)
		if err != nil {
			//newErrorResponse(c, http.StatusMovedPermanently, err.Error())
			c.Redirect(301, "/login")
			return
		}

		c.Next()
	}
}

func (h *Handler) userIdentity(c *gin.Context) {
	logrus.Info("hi")
	header := c.GetHeader(authorizationHeader)
	if header == "" {
		newErrorResponse(c, http.StatusUnauthorized, "empty auth header")
		return
	}

	headerParts := strings.Split(header, " ")
	if len(headerParts) != 2 {
		newErrorResponse(c, http.StatusUnauthorized, "invalid auth header")
		return
	}

	user, err := h.services.Authorization.ParseToken(headerParts[1])
	if err != nil {
		newErrorResponse(c, http.StatusUnauthorized, err.Error())
		return
	}

	c.Set(userCtx, user)
}

func getUserId(c *gin.Context) (int, error) {
	id, ok := c.Get(userCtx)
	if !ok {
		newErrorResponse(c, http.StatusInternalServerError, "user id not found")
		return 0, errors.New("user id not found")
	}

	userID, ok := id.(int)
	if !ok {
		newErrorResponse(c, http.StatusInternalServerError, "user id is of invalid type")
		return 0, errors.New("user id not found")
	}
	return userID, nil
}
