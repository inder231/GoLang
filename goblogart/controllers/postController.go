package controllers

import (
	"fmt"
	"goblogart/inits"
	"goblogart/models"
	"net/http"

	"github.com/gin-gonic/gin"
)

func CreatePost (ctx *gin.Context) {
	var body struct {
		Title string
		Body string
		Likes int
		Draft bool
		Author string
	}

	ctx.BindJSON(&body) // bind json of request to the struct

	post := models.Post{Title: body.Title, Body: body.Body, Likes: body.Likes, Draft: body.Draft, Author: body.Author}

	fmt.Println("----post body-----", post)
	result := inits.DB.Create(&post)

	if result.Error != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": result.Error})
		return
	}
	
	ctx.JSON(http.StatusCreated, gin.H{"data": post})
	
}

func GetPosts ( ctx *gin.Context) {
	
	var posts []models.Post
	
	result := inits.DB.Find(&posts)

	if result.Error != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": result.Error})
		return
	}

    ctx.JSON(http.StatusOK, gin.H{"data": posts})
}

func GetPost (ctx *gin.Context) {
	var post models.Post

	result := inits.DB.Find(&post, ctx.Param("id"))

	if result.Error!= nil {
        ctx.JSON(http.StatusNotFound, gin.H{"error": "Post not found"})
        return
    }

    ctx.JSON(http.StatusOK, gin.H{"data": post})
}

func UpdatePost (ctx *gin.Context) {
	var body struct {
		Title string
		Body string
		Likes int
		Draft bool
		Author string
	}

    if err := ctx.BindJSON(&body); err != nil {
        ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
        return
    }

	var post models.Post

	result := inits.DB.First(&post, ctx.Param("id"))

	if result.Error!= nil {
        ctx.JSON(http.StatusInternalServerError, gin.H{"error": result.Error})
        return
    }

	inits.DB.Model(&post).Updates(models.Post{ Title: body.Title, Body: body.Body, Likes: body.Likes, Draft: body.Draft, Author: body.Author})

	ctx.JSON(http.StatusOK, gin.H{"data": post})
}

func DeletePost (ctx *gin.Context){

	id := ctx.Param("id")

	inits.DB.Delete(&models.Post{}, id)

	ctx.JSON(http.StatusOK, gin.H{"data": "Post deleted successfully!"})
}