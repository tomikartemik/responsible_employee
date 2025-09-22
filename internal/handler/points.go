package handler

import (
    "github.com/gin-gonic/gin"
    "net/http"
    "strconv"
)

func (h *Handler) GetPointsSummary(c *gin.Context) {
    monthStr := c.Query("month")
    yearStr := c.Query("year")

    var month, year int
    var err error
    if monthStr != "" {
        month, err = strconv.Atoi(monthStr)
        if err != nil { month = 0 }
    }
    if yearStr != "" {
        year, err = strconv.Atoi(yearStr)
        if err != nil { year = 0 }
    }

    summary, err := h.services.PointsSummary(month, year)
    if err != nil {
        c.JSON(http.StatusInternalServerError, gin.H{"message": err.Error()})
        return
    }
    c.JSON(http.StatusOK, summary)
}


