package utils

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"responsible_employee/internal/model"
)

func NewErrorResponse(c *gin.Context, statusCode int, message string) {
	fmt.Println(message)
	c.AbortWithStatusJSON(statusCode, model.ErrorResponse{Message: message})
}
