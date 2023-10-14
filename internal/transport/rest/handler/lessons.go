package handler

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

func (h *Handler) AddLesson(c *gin.Context) {
	var lesson AddLesson
	if err := c.ShouldBindJSON(&lesson); err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	if err := h.services.Lesson.Add(lesson.Name, lesson.Description, lesson.OrderID, lesson.ModuleID, lesson.Content); err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.Status(http.StatusOK)
}

// func (h *Handler) AddLesson(c *gin.Context) {
// 	file, header, err := c.Request.FormFile("file")
// 	if err != nil {
// 		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
// 		return
// 	}

// 	fileName := header.Filename

// 	name := strings.Trim(c.Request.FormValue("name"), " ")
// 	description := strings.Trim(c.Request.FormValue("description"), " ")
// 	moduleName := strings.Trim(c.Request.FormValue("moduleName"), " ")
// 	courseName := strings.Trim(c.Request.FormValue("courseName"), " ")

// 	orderID, err := strconv.ParseUint(c.Request.FormValue("orderID"), 10, 0)
// 	if err != nil {
// 		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": err.Error()})
// 		return
// 	}

// 	if name == "" || description == "" || moduleName == "" || courseName == "" {
// 		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": "Bad name, description, module name or course name"})
// 		return
// 	}

// 	if err := h.services.Lesson.Add(file, fileName, name, description, uint(orderID), moduleName, courseName); err != nil {
// 		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": err.Error()})
// 		return
// 	}

// 	c.Status(http.StatusOK)
// }

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
	orderid, err := strconv.Atoi(c.Param("orderid"))

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Некорректное значение параметра orderid"})
		return
	}

	moduleid, err := strconv.Atoi(c.Param("moduleid"))

	if err != nil {
		// Обработка ошибки, если преобразование не удалось
		c.JSON(http.StatusBadRequest, gin.H{"error": "Некорректное значение параметра moduleid"})
		return
	}

	lesmd, err := h.services.Lesson.Get(moduleid, orderid)

	if err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, lesmd)
}

func (h *Handler) EditLesson(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Некорректное значение параметра id"})
		return
	}

	var edlesson EdLesson

	if err := c.ShouldBindJSON(&edlesson); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if err := h.services.Lesson.Put(id, edlesson.Name, edlesson.Description, edlesson.OrderID, edlesson.Content); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.Status(http.StatusOK)
}
