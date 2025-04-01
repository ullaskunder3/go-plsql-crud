package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/ullaskunder3/go-postgres-crud/handlers"
)

// SetupRouter initializes API routes
func SetupRouter() *gin.Engine {
	r := gin.Default()

	// User routes
	r.POST("/users", handlers.CreateUserHandler)
	r.GET("/users", handlers.GetAllUsersHandler)
	r.GET("/users/:id", handlers.GetUserHandler)
	r.PUT("/users/:id", handlers.UpdateUserHandler)
	r.DELETE("/users/:id", handlers.DeleteUserHandler)

	return r
}
