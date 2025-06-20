package handler

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func (h *Handler) UploadTaskPhoto(c *gin.Context) {
	taskID := c.Query("task_id")

	photo, err := c.FormFile("photo")
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Unable to get photo from form"})
		return
	}

	err = h.services.SaveTaskPhoto(taskID, photo)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, "Saved photo successfully")
}

func (h *Handler) UploadReportPhoto(c *gin.Context) {
	reportID := c.Query("report_id")

	photo, err := c.FormFile("photo")
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Unable to get photo from form"})
		return
	}

	err = h.services.SaveReportPhoto(reportID, photo)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, "Saved photo successfully")
}
