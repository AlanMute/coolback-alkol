package main

import (
	"github.com/KrizzMU/coolback-alkol/db"
	"github.com/KrizzMU/coolback-alkol/handlers"
	"github.com/KrizzMU/coolback-alkol/handlers/courses_handlers"
	"github.com/KrizzMU/coolback-alkol/repository"

	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default()
	repo := repository.NewRepository(db.GetConnection())

	defer repo.CloseConnection()

	router.Handle("GET", "/", handlers.TestHandler(repo))
	router.Handle("POST", "/add/course", courses_handlers.AddCourseHandler(repo))

	router.Run(":8080")
}
