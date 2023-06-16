package controllers

import (
	"go-crud/initializers"
	"go-crud/models"

	"github.com/gin-gonic/gin"
)

func PostsCreate(c *gin.Context) {
	// data from req body
	var body struct {
		Body  string
		Title string
	}
	c.Bind((&body))
	// create
	post := models.Post{Title: body.Title, Body: body.Body}

	result := initializers.DB.Create(&post)

	if result.Error != nil {
		c.Status(400)
		return
	}

	// return
	c.JSON(200, gin.H{
		"post": post,
	})
}

func PostsIndex(c *gin.Context) {
	// get
	var posts []models.Post
	initializers.DB.Find(&posts)

	// respond
	c.JSON(200, gin.H{
		"posts": posts,
	})
}

func PostShow(c *gin.Context) {
	id := c.Param("id")
	// get
	var post models.Post
	initializers.DB.First(&post, id)

	// respond
	c.JSON(200, gin.H{
		"post": post,
	})
}

func PostUpdate(c *gin.Context) {
	id := c.Param("id")

	var body struct {
		Body  string
		Title string
	}

	c.Bind(&body)

	// find post
	var post models.Post
	initializers.DB.First(&post, id)

	// update
	initializers.DB.Model(&post).Updates(models.Post{
		Title: body.Title,
		Body:  body.Body,
	})

	c.JSON(200, gin.H{
		"post": post,
	})
}

func PostDelete(c *gin.Context) {
	id := c.Param("id")

	var post models.Post
	initializers.DB.First(&post, id)

	initializers.DB.Delete(&models.Post{}, id)

	c.JSON(200, gin.H{
		"post": post,
	})
}
