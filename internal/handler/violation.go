package handler

import (
	"errors"
	"github.com/gin-gonic/gin"
	"net/http"
	"responsible_employee/internal/utils"
	"strconv"
)

func (h *Handler) GetAllViolations(c *gin.Context) {
	violations, err := h.services.GetAllViolations()

	if err != nil {
		utils.NewErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusOK, violations)
}

func (h *Handler) GetViolationByCategory(c *gin.Context) {
	category := c.Query("category")

	if category == "" {
		utils.NewErrorResponse(c, http.StatusBadRequest, errors.New("Need category!").Error())
		return
	}

	violations, err := h.services.GetViolationByCategory(category)

	if err != nil {
		utils.NewErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusOK, violations)
}

func (h *Handler) GetViolationByID(c *gin.Context) {
	idStr := c.Query("id")

	if idStr == "" {
		utils.NewErrorResponse(c, http.StatusBadRequest, errors.New("Need violation ID!").Error())
		return
	}

	id, err := strconv.Atoi(idStr)
	if err != nil {
		utils.NewErrorResponse(c, http.StatusBadRequest, errors.New("ID must be INT!").Error())
	}

	violation, err := h.services.GetViolationByID(id)

	if err != nil {
		utils.NewErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusOK, violation)
}
