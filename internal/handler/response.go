package handler

import (
	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
)

var (
	ErrNotFields    = "not all fields are filled in"
	ErrServerError  = "server error"
	ErrAccessDenied = "Access denied"
)

type errorResponse struct {
	Message string `json:"message"`
}

func newErrorResponse(c *gin.Context, statusCode int, err error, message string) {
	if err != nil {
		logrus.Error(message, "  | AND | ", err)
	}
	c.AbortWithStatusJSON(statusCode, errorResponse{message})
}
