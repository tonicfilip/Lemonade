package routes

import (
	"lemonade/handlers"

	"github.com/gin-gonic/gin"
)

func UserRoutes(superRoute *gin.RouterGroup) {

	userRouter := superRoute.Group("/users")
	{
		userRouter.GET("/health", handlers.HealthCheck)
	}

	// Health check
	// router.GET("/health", handlers.HealthCheck)

	// // User routes
	// userGroup := router.Group("/api/v1/users")
	// {
	// 	userGroup.GET("", handlers.GetUsers)
	// 	userGroup.GET("/:id", handlers.GetUser)
	// 	userGroup.POST("", handlers.CreateUser)
	// 	userGroup.PUT("/:id", handlers.UpdateUser)
	// 	userGroup.DELETE("/:id", handlers.DeleteUser)
	// }

	// // 404 handler
	// router.NoRoute(func(c *gin.Context) {
	// 	c.JSON(404, gin.H{
	// 		"status":  "error",
	// 		"code":    404,
	// 		"message": "Route not found",
	// 	})
	// })
}
