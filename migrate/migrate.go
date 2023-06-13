package main

import (
	"github.com/kuulife/go-crud/initializers"
	"github.com/kuulife/go-crud/models"
)

func init() {
	initializers.LoadEnvVariables()
	initializers.ConnectToDB()
}

func main() {
	initializers.DB.AutoMigrate(&models.Post{})
}
