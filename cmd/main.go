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

	router.Handle("GET", "/", handlers.TestHandler(repo))

	router.Run(":8080")
}
