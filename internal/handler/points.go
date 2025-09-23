package handler

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func (h *Handler) GetPointsSummary(c *gin.Context) {
	summary, err := h.services.Points.Summary()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"message": err.Error()})
		return
	}
	c.JSON(http.StatusOK, summary)
}
