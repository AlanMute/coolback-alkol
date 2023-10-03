package handler

import (
	"net/http"
	"strconv"
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

	orderID, err := strconv.ParseUint(c.Request.FormValue("orderID"), 10, 0)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if name == "" || description == "" || moduleName == "" || courseName == "" {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": "Bad name, description, module name or course name"})
		return
	}

	if err := h.services.Lesson.Add(file, fileName, name, description, uint(orderID), moduleName, courseName); err != nil {
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

func (h *Handler) SendTrialLesson(c *gin.Context) {
	var email Email

	if err := c.ShouldBindJSON(&email); err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if err := h.services.Lesson.SendTrialLesson(email.Address); err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.Status(http.StatusOK)
}
