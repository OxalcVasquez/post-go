package controller

import (
	"go-crud-2/initializers"
	"go-crud-2/models"

	"github.com/gin-gonic/gin"
)

func PostsCreate(c *gin.Context) {

	//Get data off req body
	var body struct {
		Body  string
		Title string
	}

	c.Bind(&body)
	// Create a post
	post := models.Post{Title: body.Title, Body: body.Body}

	result := initializers.DB.Create(&post) // pass pointer of data to Create

	if result.Error != nil {
		c.Status(400)
		return
	}

	// Return it
	c.JSON(201, gin.H{
		"post": post,
	})
}

func PostsIndex(c *gin.Context) {

	//Get the posts
	var posts []models.Post
	initializers.DB.Find(&posts)
	c.JSON(200, gin.H{
		"posts": posts,
	})
}

func PostsShow(c *gin.Context) {
	//Get id off url
	id := c.Param("id")
	//Get the posts
	var post models.Post
	initializers.DB.First(&post, id)
	c.JSON(200, gin.H{
		"post": post,
	})
}

func PostsUpdate(c *gin.Context) {
	//Get id off url
	id := c.Param("id")
	//Get data off req body
	var body struct {
		Body  string
		Title string
	}
	c.Bind(&body)
	//Find the post
	var post models.Post
	initializers.DB.First(&post, id)
	//Update
	initializers.DB.Model(&post).Updates(models.Post{
		Title: body.Title,
		Body:  body.Body,
	})
	// Return it
	c.JSON(200, gin.H{
		"post": post,
	})
}

func PostsDelete(c *gin.Context) {
	//Get id off url
	id := c.Param("id")
	//Get the posts
	var post models.Post
	initializers.DB.Delete(&post, id)
	c.JSON(204, gin.H{
		"post": post,
	})
}
