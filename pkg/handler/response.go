package handler

import (
	"log"

	"github.com/gin-gonic/gin"
)

type Error struct {
	Message string `json:"message"`
}

func NewErrorResponse(c *gin.Context, statusCode int, message string) {
	log.Fatalf(message)
	c.AbortWithStatusJSON(statusCode, Error{message})
}
