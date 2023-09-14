package modules_handlers

import (
	"net/http"
	"os"
	"strings"

	"github.com/KrizzMU/coolback-alkol/handlers"
	"github.com/KrizzMU/coolback-alkol/handlers/courses_handlers"
	"github.com/KrizzMU/coolback-alkol/repository"
	"github.com/gin-gonic/gin"
)

func AddModuleHandler(repo repository.Repository) gin.HandlerFunc {
	return func(c *gin.Context) {
		var info handlers.AddCourseModule

		courseFolderName := "/courses" + "" // SELECT folder_name FROM courses WHERE name = _name
		id := 0                             // SELECT id FROM courses WHERE name = _name

		path := courses_handlers.UniqueFolder("module", courseFolderName)
		slice := strings.Split(path, "/")
		folderName := "/" + slice[len(slice)-1]

		if err := os.Mkdir("./"+path, os.ModePerm); err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}

		if err := c.ShouldBindJSON(&info); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		repo.AddModule(info.Name, info.Description, id, folderName)
	}
}
