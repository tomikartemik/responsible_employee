package handler

import (
	"errors"
	"github.com/gin-gonic/gin"
	"net/http"
	"responsible_employee/internal/utils"
)

func (h *Handler) GetAllTasks(c *gin.Context) {
	tasks, err := h.services.GetAllTasks()

	if err != nil {
		utils.NewErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusOK, tasks)
}

func (h *Handler) GetTaskById(c *gin.Context) {
	id := c.Query("id")

	if id == "" {
		utils.NewErrorResponse(c, http.StatusBadRequest, errors.New("Need task ID!").Error())
		return
	}

	task, err := h.services.TaskByID(id)

	if err != nil {
		utils.NewErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusOK, task)
}
