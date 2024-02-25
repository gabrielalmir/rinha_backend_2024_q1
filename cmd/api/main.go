package main

import router "github.com/gabrielalmir/rinha_backend_2024_q1/internal/web/router"

func main() {
	println("Server is running on port 8080")
	router.SetupRoutes()
}
