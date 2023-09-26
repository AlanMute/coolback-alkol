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
	var info DeleteLesson

	if err := c.ShouldBindJSON(&info); err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	info.Name = strings.Trim(info.Name, " ")
	info.CourseName = strings.Trim(info.CourseName, " ")
	info.ModuleName = strings.Trim(info.ModuleName, " ")

	if info.Name == "" || info.CourseName == "" || info.ModuleName == "" {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": "Bad course, module or lesson name"})
		return
	}

	if err := h.services.Lesson.Delete(info.Name, info.CourseName, info.ModuleName); err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.Status(http.StatusOK)
}

func (h *Handler) GetLesson(c *gin.Context) {
	course := c.Param("coursename")
	module := c.Param("modulename")
	lesson := c.Param("lessonname")

	strFile, err := h.services.Lesson.Get(course, module, lesson)

	if err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, strFile)
}
