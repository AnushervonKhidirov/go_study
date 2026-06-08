package main

import (
	"fmt"
	"net/http"
	routes "task_tracker/internal/delivery/http"
)

func main() {
	r := routes.InitRoutes()

	err := http.ListenAndServe("localhost:8080", r)

	if err != nil {
		fmt.Println(err)
	}
}
