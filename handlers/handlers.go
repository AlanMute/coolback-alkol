package handlers

import (
	"net/http"
	"os"
	"path/filepath"
	"strconv"

	"github.com/KrizzMU/coolback-alkol/repository"
	"github.com/gin-gonic/gin"
)

func TestHandler(repo repository.Repository) gin.HandlerFunc {
	return func(c *gin.Context) {
		c.String(http.StatusOK, "success")
	}
}

func AddCourseHandler(repo repository.Repository) gin.HandlerFunc {
	return func(c *gin.Context) {
		var info AddCourseModule

		if err := c.ShouldBindJSON(&info); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		}

		folderName := uniqueFolder(info.Name, "/courses")

		repo.AddCourse(info.Name, info.Description, folderName)
		c.Status(http.StatusOK)
	}

}

func AddModuleHandler(repo repository.Repository) gin.HandlerFunc {
	return func(c *gin.Context) {
		var info AddCourseModule

		//courseFolderName :=

		//folderName := uniqueFolder(info.Name, "/courses" + )

		if err := c.ShouldBindJSON(&info); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		}

		//repo.AddModule(info.Name, info.Description, )
	}
}

func uniqueFile(name string, folder string) string {
	extension := filepath.Ext(name)
	nameWithoutExt := name[:len(name)-len(extension)]

	for i := 1; ; i++ {
		uniqueName := nameWithoutExt + "_" + strconv.Itoa(i) + extension
		filePath := filepath.Join(folder, uniqueName)
		_, err := os.Stat(filePath)
		if os.IsNotExist(err) {
			return uniqueName
		}
	}
}

func uniqueFolder(name string, folder string) string {
	for i := 1; ; i++ {
		uniqueName := name + "_" + strconv.Itoa(i)
		filePath := filepath.Join(folder, uniqueName)
		_, err := os.Stat(filePath)
		if os.IsNotExist(err) {
			return uniqueName
		}
	}
}
