package controllers

import (
	"api-crud/initializers"
	"api-crud/models"

	"github.com/gin-gonic/gin"
)

func CreatePost(c *gin.Context) {

	var createPostBody struct {
		Body  string `json:"body"`
		Title string `json:"title"`
	}

	c.BindJSON(&createPostBody)

	newPost := models.Post{Title: createPostBody.Title, Body: createPostBody.Body}
	result := initializers.DB.Create(&newPost)

	if result.Error != nil {
		c.Status(400)
		return
	}

	c.JSON(200, gin.H{
		"message": "create post",
		"results": newPost,
	})
}

func GetAllPosts(c *gin.Context) {
	var posts []models.Post
	initializers.DB.Find(&posts)

	c.JSON(200, gin.H{
		"message": "get all posts",
		"results": posts,
	})
}

func GetSinglePost(c *gin.Context) {
	id := c.Param("id")
	var post models.Post
	initializers.DB.First(&post, id)

	c.JSON(200, gin.H{
		"message": "get single post",
		"results": post,
	})
}

func UpdatePost(c *gin.Context) {

	var updatePostBody struct {
		Body  string `json:"body"`
		Title string `json:"title"`
	}
	c.BindJSON(&updatePostBody)

	id := c.Param("id")
	var post models.Post
	initializers.DB.First(&post, id)

	initializers.DB.Model(&post).Updates(models.Post{
		Title: updatePostBody.Title,
		Body:  updatePostBody.Body,
	})

	c.JSON(200, gin.H{
		"message": "update post",
		"results": post,
	})
}

func DeletePost(c *gin.Context) {
	id := c.Param("id")
	initializers.DB.Delete(&models.Post{}, id)
	c.Status(200)
}
