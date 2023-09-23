package handler

import (
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
)

func (h *Handler) AddCourse(c *gin.Context) {
	var info AddCourse

	if err := c.ShouldBindJSON(&info); err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	info.Name = strings.Trim(info.Name, " ")
	info.Description = strings.Trim(info.Description, " ")

	if info.Name == "" || info.Description == "" {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": "Bad name or description"})
		return
	}

	if err := h.services.Course.Add(info.Name, info.Description); err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.Status(http.StatusOK)
}

func (h *Handler) GetCourseByName(c *gin.Context) {
	courseName := c.Param("name")

	courses, err := h.services.Course.GetByName(courseName)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, courses)

}

func (h *Handler) GetAllCourses(c *gin.Context) {
	courses, err := h.services.Course.GetAll()
	if err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, courses)

}
