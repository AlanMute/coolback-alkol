package handler

import (
	"github.com/gin-gonic/gin"
)

func (h *Handler) AddLesson(c *gin.Context) {

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

// func uniqueFile(name string, folder string) string {
// 	extension := filepath.Ext(name)
// 	nameWithoutExt := "lesson"

// 	for i := 1; ; i++ {
// 		uniqueName := nameWithoutExt + "_" + strconv.Itoa(i) + extension
// 		filePath := filepath.Join(folder, uniqueName)
// 		_, err := os.Stat(filePath)
// 		if os.IsNotExist(err) {
// 			return folder + "/" + uniqueName
// 		}
// 	}
// }
