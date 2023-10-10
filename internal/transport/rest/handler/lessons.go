package handler

import (
	"net/http"
	"strconv"
	"strings"

	"github.com/gin-gonic/gin"
)

// Нужен свагер!
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

// Нужен свагер!
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

// @Summary GetLesson
// @Tags lesson
// @Description Get lesson by orderID
// @ID GetLesson
// @Accept  json
// @Produce  json
// @Param orderid path int true "Lesson sequence number"
// @Param moduleid path int true "Module ID"
// @Success 200 {object} core.ModLes
// @Failure 400 {string} string "error"
// @Failure 500 {string} string "error"
// @Failure default {string} error "error"
// @Router /lesson/{orderid}/{moduleid} [get]
func (h *Handler) GetLesson(c *gin.Context) {
	orderid, err := strconv.Atoi(c.Param("orderid"))

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Некорректное значение параметра orderid"})
		return
	}

	moduleid, err := strconv.Atoi(c.Param("moduleid"))

	if err != nil {
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

// @Summary EditLesson
// @Security ApiKeyAuth
// @Tags lesson
// @Description Edit lesson by ID
// @ID EditLesson
// @Param id path int true "Lesson ID"
// @Accept  json
// @Produce  json
// @Param input body EdLesson true "Edit Lesson (OrderId starts with one)"
// @Success 200
// @Failure 400 {string} string "error"
// @Failure default {string} error "error"
// @Router /adm/lesson/{id} [put]
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
