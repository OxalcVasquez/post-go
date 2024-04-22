package main

import (
	"go-crud-2/initializers"
	"go-crud-2/models"
)

func init() {
	initializers.LoadAndVariables()
	initializers.ConnectToDb()
}

func main() {
	initializers.DB.AutoMigrate(&models.Post{})
}
