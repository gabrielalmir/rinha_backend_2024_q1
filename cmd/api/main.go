package main

import "github.com/gabrielalmir/rinha_backend_2024_q1/internal/router"

func main() {
	println("Server is running on port 8080")
	router.SetupRoutes()
}
