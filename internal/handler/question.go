package handler

import (
	"errors"
	"github.com/gin-gonic/gin"
	"net/http"
	"responsible_employee/internal/utils"
	"strconv"
)

func (h *Handler) GetQuestionById(c *gin.Context) {
	idStr := c.Query("id")

	if idStr == "" {
		utils.NewErrorResponse(c, http.StatusBadRequest, errors.New("Need question ID!").Error())
		return
	}

	id, err := strconv.Atoi(idStr)

	if err != nil {
		utils.NewErrorResponse(c, http.StatusBadRequest, errors.New("Need question ID!").Error())
		return
	}

	report, err := h.services.QuestionByID(id)

	if err != nil {
		utils.NewErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusOK, report)
}
