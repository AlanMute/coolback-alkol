package handler

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

// @Summary AddLesson
// @Security ApiKeyAuth
// @Tags lesson
// @Description Add lesson
// @ID AddLesson
// @Accept  json
// @Produce  json
// @Param input body AddLesson true "Add Lesson (OrderId starts with one)"
// @Success 200
// @Failure 400 {string} string "error"
// @Failure default {string} error "error"
// @Router /adm/lesson/ [post]
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

// @Summary DeleteLesson
// @Security ApiKeyAuth
// @Tags lesson
// @Description Delete lesson by ID
// @ID DeleteLesson
// @Accept  json
// @Produce  json
// @Param input body Delete true "Lesson ID"
// @Success 200
// @Failure 400 {string} string "error"
// @Failure default {string} error "error"
// @Router /adm/lesson/ [DELETE]
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

// @Summary SendLetter
// @Tags email
// @Description Send email trial lesson
// @ID SendTrialLesson
// @Accept  json
// @Produce  json
// @Param input body Email true "Recipient's e-mail"
// @Success 200
// @Failure 400 {string} string "error"
// @Failure default {string} error "error"
// @Router /lesson/trial [post]
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
