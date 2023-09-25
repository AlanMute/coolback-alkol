package handler

import (
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
)

func (h *Handler) AddModule(c *gin.Context) {
	var info AddModule

	if err := c.ShouldBindJSON(&info); err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	info.Name = strings.Trim(info.Name, " ")
	info.Description = strings.Trim(info.Description, " ")
	info.CourseName = strings.Trim(info.CourseName, " ")

	if info.Name == "" || info.Description == "" || info.CourseName == "" {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": "Bad name, description or course name"})
		return
	}

	if err := h.services.Module.Add(info.Name, info.Description, info.CourseName); err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.Status(http.StatusOK)
}

func (h *Handler) DeleteModule(c *gin.Context) {

}
