package handler

import (
	"errors"
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
)

const (
	authorizationHeader = "Authorization"
	userCtx             = "userID"
)

func (h *Handler) userIdentity(c *gin.Context) {
	header := c.GetHeader(authorizationHeader)
	if header == "" {
		newErrorResponse(c, http.StatusUnauthorized, errors.New("EPMTY AUTH HEADER"), "empty auth header")
		return
	}

	headerParts := strings.Split(header, " ")
	if len(headerParts) != 2 {
		newErrorResponse(c, http.StatusUnauthorized, errors.New("INVALID AUTH HEADER"), "invalid auth header")
		return
	}

	user, err := h.services.Authorization.ParseToken(headerParts[1])
	if err != nil {
		newErrorResponse(c, http.StatusUnauthorized, err, "unauthorized")
		return
	}

	userInfo, err := h.services.Users.GetUser(user)
	if err != nil {
		newErrorResponse(c, http.StatusUnauthorized, err, "userInfo error")
		return
	}

	if userInfo.Banned == 1 || userInfo.IsDeleted == 1 {
		newErrorResponse(c, http.StatusUnauthorized, err, "block or delete")
		return
	}

	c.Set(userCtx, user)
}

func getUserId(c *gin.Context) (int, error) {
	id, ok := c.Get(userCtx)
	if !ok {
		newErrorResponse(c, http.StatusInternalServerError, errors.New("TOKEN - USER ID NOT FOUND"), "user id not found")
		return 0, errors.New("user id not found")
	}

	userID, ok := id.(int)
	if !ok {
		newErrorResponse(c, http.StatusInternalServerError, errors.New("TOKEN - USER ID INVALID TYPE"), "user id is of invalid type")
		return 0, errors.New("user id not found")
	}
	return userID, nil
}
