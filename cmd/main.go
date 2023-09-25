package main

import (
	"log"

	"github.com/KrizzMU/coolback-alkol/internal/db"
	"github.com/KrizzMU/coolback-alkol/internal/repository"
	"github.com/KrizzMU/coolback-alkol/internal/service"
	"github.com/KrizzMU/coolback-alkol/internal/transport/rest"
	"github.com/KrizzMU/coolback-alkol/internal/transport/rest/handler"
	//"github.com/gin-gonic/gin"
)

func main() {
	s := new(rest.Server)

	repos := repository.NewRepository(db.GetConnection())
	services := service.NewService(repos)
	handlers := handler.NewHandler(services)

	if err := s.Run("8080", handlers.InitRoutes()); err != nil {
		log.Fatal("ERROR start Server!")
	}
}
