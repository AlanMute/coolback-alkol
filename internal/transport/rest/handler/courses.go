package handler

import (
	"net/http"
	"strconv"
	"strings"

	"github.com/gin-gonic/gin"
)

// @Summary AddCourse
// @Security ApiKeyAuth
// @Tags course
// @Description Add Course
// @ID AddCourse
// @Accept  json
// @Produce  json
// @Param input body AddCourse true "Add course"
// @Success 200
// @Failure 400 {string} string "error"
// @Failure default {string} error "error"
// @Router /adm/course/ [post]
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

// @Summary DeleteCourse
// @Security ApiKeyAuth
// @Tags course
// @Description Delete Course by ID
// @ID DeleteCourse
// @Accept  json
// @Produce  json
// @Param input body Delete true "Course ID"
// @Success 200
// @Failure 400 {string} string "error"
// @Failure default {string} error "error"
// @Router /adm/course/ [DELETE]
func (h *Handler) DeleteCourse(c *gin.Context) {
	var info Delete

	if err := c.ShouldBindJSON(&info); err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	if err := h.services.Course.Delete(info.ID); err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.Status(http.StatusOK)
}

// @Summary Searching Courses
// @Tags course
// @Description Get courses by name
// @ID GetCourseByName
// @Accept  json
// @Produce  json
// @Param name path string true "Имя курса для поиска"
// @Success 200 {object} core.Course
// @Failure 500 {string} string "error"
// @Failure default {string} error "error"
// @Router /course/search/{name} [get]
func (h *Handler) GetCourseByName(c *gin.Context) {
	courseName := c.Param("name")

	courses, err := h.services.Course.GetByName(courseName)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, courses)

}

// @Summary All courses
// @Tags course
// @Description Get all courses
// @ID GetAllCourses
// @Accept  json
// @Produce  json
// @Success 200 {object} core.Course
// @Failure 500 {string} string "error"
// @Failure default {string} error "error"
// @Router /course/getall [get]
func (h *Handler) GetAllCourses(c *gin.Context) {
	courses, err := h.services.Course.GetAll()
	if err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, courses)

}

// @Summary GetCourse
// @Tags course
// @Description Get courses by ID with all modules&lessons
// @ID GetCourse
// @Accept  json
// @Produce  json
// @Param id path int true "Course ID"
// @Success 200 {object} core.Course
// @Failure 400 {string} string "error"
// @Failure 500 {string} string "error"
// @Failure default {string} error "error"
// @Router /course/{id} [get]
func (h *Handler) GetCourse(c *gin.Context) {

	id, err := strconv.Atoi(c.Param("id"))

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Некорректное значение параметра id"})
		return
	}

	content, err := h.services.Course.Get(id)

	if err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, content)

}

// @Summary EditCourse
// @Security ApiKeyAuth
// @Tags course
// @Description Edit course by ID
// @ID EditCourse
// @Param id path int true "Идентификатор курса для обновления"
// @Accept  json
// @Produce  json
// @Param input body AddCourse true "Edit Course"
// @Success 200
// @Failure 400 {string} string "error"
// @Failure 500 {string} string "error"
// @Failure default {string} error "error"
// @Router /adm/course/{id} [put]
func (h *Handler) EditCourse(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Некорректное значение параметра id"})
		return
	}

	var edcourse AddCourse

	if err := c.ShouldBindJSON(&edcourse); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if err := h.services.Course.Put(id, edcourse.Name, edcourse.Description); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.Status(http.StatusOK)
}

// @Summary Upload Course Image
// @Security ApiKeyAuth
// @Description Uploads an image for a course by its ID. Also can be used to change image.
// @Tags course
// @ID UploadCourseImage
// @Accept multipart/form-data
// @Produce json
// @Param id path int true "Course ID"
// @Param image formData file true "Image file to upload for the course."
// @Success 200 {string} string "OK"
// @Failure 400 {string} string "Bad Request"
// @Failure 404 {string} string "Not Found"
// @Failure 500 {string} string "Internal Server Error"
// @Router /adm/course/image/{id} [post]
func (h *Handler) UploadCourseImage(c *gin.Context) {
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

	filepath, err := h.services.Course.GetImage(uint(id))
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

// @Summary Download Course Image
// @Description Downloads an image for a course by its ID.
// @Tags course
// @ID DownloadCourseImage
// @Produce image/*
// @Param id path int true "Course ID"
// @Success 200 {string} string "OK"
// @Failure 400 {string} string "Bad Request"
// @Failure 404 {string} string "Not Found"
// @Failure 500 {string} string "Internal Server Error"
// @Router /adm/course/image/{id} [get]
func (h *Handler) DownloadCourseImage(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": err.Error()})
	}

	filepath, err := h.services.Course.GetImage(uint(id))
	if err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.File(filepath)
}

// @Summary Delete Course Image
// @Security ApiKeyAuth
// @Description Deletes an image for a course by its ID.
// @Tags course
// @ID DeleteCourseImage
// @Produce json
// @Param id path int true "Course ID"
// @Success 200 {string} string "OK"
// @Failure 400 {string} string "Bad Request"
// @Failure 404 {string} string "Not Found"
// @Failure 500 {string} string "Internal Server Error"
// @Router /adm/course/image/{id} [delete]
func (h *Handler) DeleteCourseImage(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if err := h.services.Course.DeleteImage(uint(id)); err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.Status(http.StatusOK)
}
