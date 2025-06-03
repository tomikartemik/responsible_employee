package handler

import (
	"errors"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
	"net/http"
	"responsible_employee/internal/model"
	"responsible_employee/internal/utils"
)

func (h *Handler) SignUp(c *gin.Context) {
	var user model.User

	if err := c.ShouldBindJSON(&user); err != nil {
		utils.NewErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}

	if err := h.services.SignUp(user); err != nil {
		if errors.Is(err, gorm.ErrDuplicatedKey) {
			utils.NewErrorResponse(c, http.StatusConflict, "Пользователь с таким никнеймом уже существует!")
			return
		}
		utils.NewErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusOK, model.SuccessResponse{Message: "Пользователь успешно зарегистрирован!"})
}

func (h *Handler) SignIn(c *gin.Context) {
	var input model.SignInInput

	if err := c.ShouldBindJSON(&input); err != nil {
		utils.NewErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}

	user, err := h.services.SignIn(input)
	if err != nil {
		if err.Error() == "Пользователя с таким никнеймом не существует!" {
			utils.NewErrorResponse(c, http.StatusUnauthorized, err.Error())
			return
		} else if err.Error() == "Неверный пароль!" {
			utils.NewErrorResponse(c, http.StatusUnauthorized, err.Error())
			return
		}
		utils.NewErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusOK, user)
}

func (h *Handler) GetUsersSortedByPoints(c *gin.Context) {
	users, err := h.services.GetUsersSortedByPoints()

	if err != nil {
		utils.NewErrorResponse(c, http.StatusInternalServerError, err.Error())
	}

	c.JSON(http.StatusOK, users)
}

func (h *Handler) CreateTask(c *gin.Context) {
	var task model.Task

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

	if err := c.ShouldBindJSON(&task); err != nil {
		utils.NewErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}

	if err := h.services.CreateTask(task, userID); err != nil {
		utils.NewErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusOK, model.SuccessResponse{Message: "Проблема успешно сохранена!"})
}

func (h *Handler) TakeTask(c *gin.Context) {
	taskID := c.Query("taskId")

	if taskID == "" {
		utils.NewErrorResponse(c, http.StatusBadRequest, "Task ID is required")
		return
	}

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

	err := h.services.TakeTask(userID, taskID)

	if err != nil {
		utils.NewErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusOK, "Task taken successfully")
}

func (h *Handler) CompleteTask(c *gin.Context) {

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

	var report model.Report

	if err := c.ShouldBindJSON(&report); err != nil {
		utils.NewErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}

	report.UserID = userID

	if err := h.services.RegisterReport(report); err != nil {
		utils.NewErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusOK, "Task completed successfully")
}
