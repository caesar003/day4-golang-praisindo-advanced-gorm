package router

import (
	"github.com/caesar003/day4-golang-praisindo-advanced-gorm/handlers"
	"github.com/caesar003/day4-golang-praisindo-advanced-gorm/middleware"

	"github.com/gin-gonic/gin"
)

func SetupRouter() *gin.Engine {
	r := gin.Default()

	r.GET("/", handlers.RootHandler)

	publicRoutes := r.Group("/api/user")
	privateRoutes := r.Group("/api/user")

	privateRoutes.Use(middleware.AuthMiddleWare())

	publicRoutes.GET("/", handlers.GetUsers)
	publicRoutes.GET("/:id", handlers.GetUserByID)

	privateRoutes.POST("/", handlers.CreateUser)
	privateRoutes.PUT("/:id", handlers.UpdateUser)
	privateRoutes.DELETE("/:id", handlers.DeleteUser)

	return r
}
