package main

import (
	"fmt"
	"log"
	"net/http"
	"task_tracker/internal/config"
	"task_tracker/internal/database"
	routes "task_tracker/internal/delivery/http"
	"task_tracker/internal/handler"
	"task_tracker/internal/repository"
	"task_tracker/internal/service"
)

func main() {
	cfg := config.InitConfig()

	db, err := database.Connect(cfg)

	if err != nil {
		log.Fatalln(err)
	}

	repos := repository.NewRepository(db)
	services := service.NewService(repos)
	handlers := handler.NewHandler(services)

	r := routes.InitRoutes(handlers)
	err = http.ListenAndServe(cfg.Address, r)

	if err != nil {
		fmt.Println(err)
	}
}
