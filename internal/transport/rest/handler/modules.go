package handler

import (
	"net/http"
	"strconv"
	"strings"

	"github.com/gin-gonic/gin"
)

// Нужен свагер!
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

// Нужен свагер!
func (h *Handler) DeleteModule(c *gin.Context) {
	var info Delete

	if err := c.ShouldBindJSON(&info); err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	if err := h.services.Module.Delete(info.ID); err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.Status(http.StatusOK)
}

// @Summary GetModule
// @Tags module
// @Description Get module by ID
// @ID GetModule
// @Accept  json
// @Produce  json
// @Param id path int true "Module ID"
// @Success 200 {object} core.ModLes
// @Failure 400 {string} string "error"
// @Failure 500 {string} string "error"
// @Failure default {string} error "error"
// @Router /module/{id} [get]
func (h *Handler) GetModule(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Некорректное значение параметра id"})
		return
	}

	modles, err := h.services.Module.Get(id)

	if err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, modles)
}

func (h *Handler) EditModule(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Некорректное значение параметра id"})
		return
	}

	var edmodule EdModule
	if err := c.ShouldBindJSON(&edmodule); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if err := h.services.Module.Put(id, edmodule.Name, edmodule.Description, edmodule.OrderID); err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.Status(http.StatusOK)

}
