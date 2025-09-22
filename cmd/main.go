package main

import (
	"github.com/joho/godotenv"
	"github.com/sirupsen/logrus"
	"log"
	"os"
	"responsible_employee/internal"
	"responsible_employee/internal/handler"
	"responsible_employee/internal/repository"
	"responsible_employee/internal/service"
	"responsible_employee/internal/utils"
)

func main() {
	if err := godotenv.Load(); err != nil {
		log.Println("Warning: No .env file found, using system environment variables.")
	}

	db, err := repository.ConnectDB()
	if err != nil {
		log.Fatalf("Failed to connect to the database: %v", err)
	}

	log.Println("Application started successfully")

	repos := repository.NewRepository(db)
	services := service.NewService(repos)
	handlers := handler.NewHandler(services)

	taskChecker := utils.NewTaskChecker(db)
	taskChecker.Start()
	defer taskChecker.Stop()

	monthlyResetter := utils.NewMonthlyResetter(db)
	monthlyResetter.Start()
	defer monthlyResetter.Stop()

	srv := new(internal.Server)
	if err := srv.Run(os.Getenv("PORT"), handlers.InitRoutes()); err != nil {
		logrus.Fatalf("error occured while running server %s", err.Error())
	}

}
