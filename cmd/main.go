package main

import (
	"github.com/KrizzMU/coolback-alkol/db"
	"github.com/KrizzMU/coolback-alkol/handlers"
	"github.com/KrizzMU/coolback-alkol/repository"

	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default()
	repo := repository.NewRepository(db.GetConnection())

	defer repo.CloseConnection()

	router.Handle("GET", "/", handlers.TestHandler(repo))
	router.Handle("POST", "/add/course", handlers.AddCourseHandler(repo))
	router.Handle("POST", "/add/module", handlers.AddModuleHandler(repo))
	router.Handle("POST", "/add/lesson", handlers.AddLessonHandler(repo))

	router.Run(":8080")
}
