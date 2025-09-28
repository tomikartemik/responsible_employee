package handler

import (
	"errors"
	"github.com/gin-gonic/gin"
	"net/http"
	"responsible_employee/internal/model"
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

func (h *Handler) GenerateTest(c *gin.Context) {
	category := c.Query("category")
	test, err := h.services.GenerateTest(category)
	if err != nil {
		utils.NewErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusOK, test)
}

func (h *Handler) CheckAnswers(c *gin.Context) {
	userIDStr, exists := c.Get("user_id")
	if !exists {
		utils.NewErrorResponse(c, http.StatusUnauthorized, "user_id not found in context")
		return
	}

	userID, ok := userIDStr.(string)
	if !ok {
		utils.NewErrorResponse(c, http.StatusUnauthorized, "invalid user_id type in context")
		return
	}

	var answers model.TestInput

	if err := c.ShouldBindJSON(&answers); err != nil {
		utils.NewErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}

	result, err := h.services.CheckUserAnswers(userID, answers)
	if err != nil {
		utils.NewErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusOK, result)
}
