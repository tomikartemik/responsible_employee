package handler

import (
	"errors"
	"github.com/gin-gonic/gin"
	"net/http"
	"responsible_employee/internal/service"
	"responsible_employee/internal/utils"
	"strings"
)

const (
	authorizationHeader = "Authorization"
)

func (h *Handler) UserIdentity(c *gin.Context) {
	header := c.GetHeader(authorizationHeader)
	if header == "" {
		utils.NewErrorResponse(c, http.StatusUnauthorized, "No authorization header")
		return
	}

	token, err := extractToken(header)
	if err != nil {
		utils.NewErrorResponse(c, http.StatusUnauthorized, err.Error())
		return
	}

	userId, err := service.ParseToken(token)
	if err != nil {
		utils.NewErrorResponse(c, http.StatusUnauthorized, err.Error())
		return
	}

	c.Set("user_id", userId)
}

func extractToken(header string) (string, error) {
	if header == "" {
		return "", errors.New("empty auth header")
	}

	parts := strings.Split(header, " ")
	if len(parts) != 2 || parts[0] != "Bearer" {
		return "", errors.New("invalid auth header")
	}

	return parts[1], nil
}
