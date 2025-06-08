package handler

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"responsible_employee/internal/utils"
	"strconv"
)

func (h *Handler) GetMessageByUserID(c *gin.Context) {

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

	messages, err := h.services.MessagesByUserID(userID)

	if err != nil {
		utils.NewErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusOK, messages)
}

func (h *Handler) ReadMessageByID(c *gin.Context) {
	messageIDStr := c.Query("message_id")

	messageID, err := strconv.Atoi(messageIDStr)
	if err != nil {
		utils.NewErrorResponse(c, http.StatusBadRequest, "invalid message id")
		return
	}

	err = h.services.ReadMessage(messageID)
	if err != nil {
		utils.NewErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusOK, "Successfully read message")
}
