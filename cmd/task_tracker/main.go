package main

import (
	"fmt"
	"net/http"
	"task_tracker/internal/config"
	routes "task_tracker/internal/delivery/http"
)

func main() {
	cfg := config.InitConfig()

	r := routes.InitRoutes()

	err := http.ListenAndServe(cfg.Address, r)

	if err != nil {
		fmt.Println(err)
	}
}
