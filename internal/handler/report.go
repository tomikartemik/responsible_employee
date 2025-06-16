package handler

import (
	"errors"
	"github.com/gin-gonic/gin"
	"net/http"
	"responsible_employee/internal/utils"
)

func (h *Handler) GetReportById(c *gin.Context) {
	id := c.Query("id")

	if id == "" {
		utils.NewErrorResponse(c, http.StatusBadRequest, errors.New("Need report ID!").Error())
		return
	}

	report, err := h.services.ReportByID(id)

	if err != nil {
		utils.NewErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusOK, report)
}
