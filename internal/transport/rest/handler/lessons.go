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

func (h *Handler) GetLesson(c *gin.Context) {

}

// Это должно быть реализовано в репозитории, тк тут идет работа с данными
// func  AddLessonHandler(repo repository.Repository) gin.HandlerFunc {
// 	return func(c *gin.Context) {
// 		file, header, err := c.Request.FormFile("file")
// 		if err != nil {
// 			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
// 			return
// 		}

// 		ext := filepath.Ext(header.Filename)

// 		moduleFolderName := "/courses" + "" + ""
// 		// SELECT folder_name FROM courses WHERE id = (SELECT course_id FROM modules WHERE name = _name)
// 		// SELECT folder_name FROM modules WHERE name = _name
// 		id := 0 // SELECT id FROM modules WHERE name = _name

// 		path := uniqueFile("lesson."+ext, moduleFolderName)
// 		slice := strings.Split(path, "/")
// 		fileName := "/" + slice[len(slice)-1]

// 		out, err := os.Create(path)
// 		if err != nil {
// 			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
// 			return
// 		}

// 		if _, err = io.Copy(out, file); err != nil {
// 			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
// 			return
// 		}

// 		if err = file.Close(); err != nil {
// 			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
// 			return
// 		}

// 		if err = out.Close(); err != nil {
// 			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
// 			return
// 		}

// 		name := c.Request.FormValue("name")
// 		description := c.Request.FormValue("description")

// 		if err = repo.AddLesson(name, description, id, fileName); err != nil {
// 			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
// 			return
// 		}
// 	}
// }
