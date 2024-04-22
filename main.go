package main

import (
	"go-crud-2/controller"
	"go-crud-2/initializers"
	"go-crud-2/models"

	"github.com/gin-gonic/gin"
)

func init() {
	initializers.LoadAndVariables()
	initializers.ConnectToDb()

}

func main() {

	r := gin.Default() //app router
	initializers.DB.AutoMigrate(&models.Post{})

	r.POST("posts", controller.PostsCreate)
	r.GET("posts", controller.PostsIndex)
	r.GET("posts/:id", controller.PostsShow)
	r.PUT("posts/:id", controller.PostsUpdate)
	r.DELETE("posts/:id", controller.PostsDelete)

	r.Run()
}
