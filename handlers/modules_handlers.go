package handlers

import (
	"net/http"
	"os"
	"strings"

	"github.com/KrizzMU/coolback-alkol/repository"
	"github.com/gin-gonic/gin"
)

func AddModuleHandler(repo repository.Repository) gin.HandlerFunc {
	return func(c *gin.Context) {
		var info AddCourseModule

		courseFolderName := "/courses" + "" // SELECT folder_name FROM courses WHERE name = _name
		id := 0                             // SELECT id FROM courses WHERE name = _name

		path := uniqueFolder("module", courseFolderName)
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
