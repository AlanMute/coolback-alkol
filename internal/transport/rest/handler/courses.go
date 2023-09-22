package handler

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func (h *Handler) AddCourse(c *gin.Context) {

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

// Входные данные тут не нужны так как все хранится в контексте будет

// func AddCourse(repo repository.Repository) gin.HandlerFunc {
// 	return func(c *gin.Context) {
// 		var info AddCourseModule

// 		if err := c.ShouldBindJSON(&info); err != nil {
// 			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
// 			return
// 		}

// 		path := uniqueFolder("course", "/courses")
// 		slice := strings.Split(path, "/")
// 		folderName := "/" + slice[len(slice)-1]

// 		if err := os.Mkdir("./"+path, os.ModePerm); err != nil {
// 			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
// 			return
// 		}

// 		repo.AddCourse(info.Name, info.Description, folderName)
// 		c.Status(http.StatusOK)
// 	}
// }

// func uniqueFolder(name string, folder string) string {
// 	for i := 1; ; i++ {
// 		uniqueName := name + "_" + strconv.Itoa(i)
// 		filePath := filepath.Join(folder, uniqueName)
// 		_, err := os.Stat(filePath)
// 		if os.IsNotExist(err) {
// 			return folder + "/" + uniqueName
// 		}
// 	}
// }
