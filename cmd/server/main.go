package main

import "github.com/xuanvinh9411/go-ecommerce-backend/internal/routers"


func main() {
	// Create a new Gin router
	r := routers.NewRouter()
	
	r.Run(":3000") // Start the server on port 3000
	
}

