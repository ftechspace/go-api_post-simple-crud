package main

import (
	"api-crud/controllers"
	"api-crud/initializers"

	// "gorm.io/gorm"
	// "gorm.io/driver/sqlite"
	"github.com/gin-gonic/gin"
)

func init() {
	initializers.LoadEnvVariables()
	initializers.ConnectToDB()
}

func main() {
	r := gin.Default()

	r.POST("/posts/", controllers.CreatePost)
	r.GET("/posts", controllers.GetAllPosts)
	r.GET("/posts/:id", controllers.GetSinglePost)
	r.PUT("/posts/:id", controllers.UpdatePost)
	r.DELETE("/posts/:id", controllers.DeletePost)

	r.Run()
}
