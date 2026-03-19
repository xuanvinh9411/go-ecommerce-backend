package routers

import (
	"github.com/gin-gonic/gin"
	"github.com/xuanvinh9411/go-ecommerce-backend/internal/controller"
	"github.com/xuanvinh9411/go-ecommerce-backend/internal/middlewares"
	"github.com/xuanvinh9411/go-ecommerce-backend/pkg/logger"
)

func NewRouter() *gin.Engine {
	// Create a new Gin router
	// r := gin.Default()
	r := gin.New()
	r.Use(middlewares.LoggerMiddleware((logger.GetLogger()))) // Use the custom logger middleware
	v1 :=r.Group("/v1")
	v1.GET("/users", controller.NewUserController().GetUser) // Define a route for GET /users	
	v1.GET("/test-create-user/:uid", controller.NewUserController().CreateUser) // Define a route for GET /users	
	v1.GET("/query-user", controller.NewUserController().GetUserQuery) // Define a route for GET /users	
	r.Run(":3000") // Start the server on port 3000
	
	return r
}

