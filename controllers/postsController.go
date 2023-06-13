package controllers

import (
	"github.com/gin-gonic/gin"
	"github.com/kuulife/go-crud/initializers"
	"github.com/kuulife/go-crud/models"
)

func PostsCreate(c *gin.Context) {
	//get data off req body
	var body struct {
		Body  string
		Title string
	}

	c.Bind(&body)

	//create a post
	post := models.Post{Title: body.Title, Body: body.Body}

	result := initializers.DB.Create(&post)
	if result.Error != nil {
		c.Status(400)
		return
	}
	//return it
	c.JSON(200, gin.H{
		"post": post,
	})
}

func PostsIndex(c *gin.Context) {
	//get the posts
	var posts []models.Post
	initializers.DB.Find(&posts)

	//respond with them
	c.JSON(200, gin.H{
		"posts": posts,
	})
}

func PostsShow(c *gin.Context) {

	//get url
	id := c.Param("id")
	//get the post
	var post models.Post
	initializers.DB.Find(&post, id)

	//respond with post
	c.JSON(200, gin.H{
		"post": post,
	})
}

func PostUpdate(c *gin.Context) {
	//get the id of the url
	id := c.Param("id")

	//get the req of the body
	var body struct {
		Body  string
		Title string
	}

	c.Bind(&body)

	//fid the post for updating
	var post models.Post
	initializers.DB.Find(&post, id)

	//update
	initializers.DB.Model(&post).Updates(models.Post{Title: body.Title, Body: body.Body})

	//respond with it

	c.JSON(200, gin.H{
		"post": post,
	})
}

func PostDelete(c *gin.Context) {
	//find the id
	id := c.Param("id")

	//delete the post
	initializers.DB.Delete(&models.Post{}, id)

	//respond
	c.Status(200)
}
