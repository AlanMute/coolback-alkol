package handler

import (
	"net/http"
	"strconv"
	"strings"

	"github.com/gin-gonic/gin"
)

// @Summary AddModule
// @Security ApiKeyAuth
// @Tags module
// @Description Add Module
// @ID AddModule
// @Accept  json
// @Produce  json
// @Param input body AddModule true "Add Module (OrderId starts with one)"
// @Success 200
// @Failure 400 {string} string "error"
// @Failure default {string} error "error"
// @Router /adm/module/ [post]
func (h *Handler) AddModule(c *gin.Context) {
	var info AddModule

	if err := c.ShouldBindJSON(&info); err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	info.Name = strings.Trim(info.Name, " ")
	info.Description = strings.Trim(info.Description, " ")

	if err := h.services.Module.Add(info.Name, info.Description, info.OrderID, info.CourseID); err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.Status(http.StatusOK)
}

// @Summary DeleteModule
// @Security ApiKeyAuth
// @Tags module
// @Description Delete Module by ID
// @ID DeleteModule
// @Accept  json
// @Produce  json
// @Param input body Delete true "Module ID"
// @Success 200
// @Failure 400 {string} string "error"
// @Failure default {string} error "error"
// @Router /adm/module/ [DELETE]
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

// @Summary EditModule
// @Security ApiKeyAuth
// @Tags module
// @Description Edit module by ID
// @ID EditModule
// @Param id path int true "Module ID"
// @Accept  json
// @Produce  json
// @Param input body EdModule true "Edit Module (OrderId starts with one)"
// @Success 200
// @Failure 400 {string} string "error"
// @Failure default {string} error "error"
// @Router /adm/module/{id} [put]
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

// @Summary Upload Module Image
// @Security ApiKeyAuth
// @Description Uploads an image for a module by its ID. Also can be used to change image.
// @Tags module
// @ID UploadModuleImage
// @Accept multipart/form-data
// @Produce json
// @Param id path int true "Module ID"
// @Param image formData file true "Image file to upload for the module."
// @Success 200 {string} string "OK"
// @Failure 400 {string} string "Bad Request"
// @Failure 404 {string} string "Not Found"
// @Failure 500 {string} string "Internal Server Error"
// @Router /adm/module/image/{id} [post]
func (h *Handler) UploadModuleImage(c *gin.Context) {
	file, err := c.FormFile("image")
	if err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	filepath, err := h.services.Module.GetImage(uint(id))
	if err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	if err := c.SaveUploadedFile(file, filepath); err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.Status(http.StatusOK)
}

// @Summary Download Module Image
// @Description Downloads an image for a module by its ID.
// @Tags module
// @ID DownloadModuleImage
// @Produce image/*
// @Param id path int true "Module ID"
// @Success 200 {string} string "OK"
// @Failure 400 {string} string "Bad Request"
// @Failure 404 {string} string "Not Found"
// @Failure 500 {string} string "Internal Server Error"
// @Router /adm/module/image/{id} [get]
func (h *Handler) DownloadModuleImage(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	filepath, err := h.services.Module.GetImage(uint(id))
	if err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.File(filepath)
}

// @Summary Delete Module Image
// @Security ApiKeyAuth
// @Description Deletes an image for a module by its ID.
// @Tags module
// @ID DeleteModuleImage
// @Produce json
// @Param id path int true "Module ID"
// @Success 200 {string} string "OK"
// @Failure 400 {string} string "Bad Request"
// @Failure 404 {string} string "Not Found"
// @Failure 500 {string} string "Internal Server Error"
// @Router /adm/module/image/{id} [delete]
func (h *Handler) DeleteModuleImage(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if err := h.services.Module.DeleteImage(uint(id)); err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.Status(http.StatusOK)
}
