package main

import (
	"log"
	"os"

	"github.com/KrizzMU/coolback-alkol/internal/db"
	"github.com/KrizzMU/coolback-alkol/internal/repository"
	"github.com/KrizzMU/coolback-alkol/internal/service"
	"github.com/KrizzMU/coolback-alkol/internal/transport/rest"
	"github.com/KrizzMU/coolback-alkol/internal/transport/rest/handler"
	"github.com/KrizzMU/coolback-alkol/pkg/auth"
	"github.com/joho/godotenv"
)

//@title eptanit.com
//@version 1.1
//@description EptaNit Super Site
//contact.email eptanit@gmail.com

// @host localhost:8080
// @BasePath /

// @securityDefinitions.apikey ApiKeyAuth
// @in header
// @name Authorization

func main() {
	s := new(rest.Server)

	if err := godotenv.Load(); err != nil {
		panic(err)
	}

	tokenManager, err := auth.NewManager(os.Getenv("JWT_SIGNING_KEY"))
	if err != nil {
		log.Fatal("Fatal tokenManager:", err.Error())
	}
	repos := repository.NewRepository(db.GetConnection())
	services := service.NewService(repos, tokenManager)
	handlers := handler.NewHandler(services, tokenManager)

	if err := s.Run("8080", handlers.InitRoutes()); err != nil {
		log.Fatal("ERROR start Server!")
	}
}
