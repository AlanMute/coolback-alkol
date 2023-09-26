package handler

import (
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
)

func (h *Handler) AddLesson(c *gin.Context) {
	file, header, err := c.Request.FormFile("file")
	if err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	fileName := header.Filename

	name := strings.Trim(c.Request.FormValue("name"), " ")
	description := strings.Trim(c.Request.FormValue("description"), " ")
	moduleName := strings.Trim(c.Request.FormValue("moduleName"), " ")
	courseName := strings.Trim(c.Request.FormValue("courseName"), " ")

	if name == "" || description == "" || moduleName == "" || courseName == "" {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": "Bad name, description, module name or course name"})
		return
	}

	if err := h.services.Lesson.Add(file, fileName, name, description, moduleName, courseName); err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.Status(http.StatusOK)
}

func (h *Handler) DeleteLesson(c *gin.Context) {
	var info Delete

	if err := c.ShouldBindJSON(&info); err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	if err := h.services.Lesson.Delete(info.ID); err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.Status(http.StatusOK)
}

func (h *Handler) GetLesson(c *gin.Context) {

}
