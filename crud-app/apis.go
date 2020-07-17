package main

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func getAllPosts(g *gin.Context) {
	// query parameters
	limit := g.DefaultQuery("limit", "10")
	offset := g.DefaultQuery("offset", "0")

	var posts []Post
	db.Limit(limit).Offset(offset).Find(&posts)
	g.JSON(http.StatusOK, gin.H{
		"error":"",
		"data": posts,
	})
}

func getSinglePost(g *gin.Context) {
	post := getPostById(g)
	if post.ID == 0 {
		return
	}
	g.JSON(http.StatusOK, gin.H{
		"error":"",
		"data": post,
	})
}

func createPost(g *gin.Context) {
	// instance from the Post; which is the model and the validator in the same time
	var post Post
	// if there is an error; the sent data doesnt match with the validations
	if err := g.ShouldBindJSON(&post); err != nil {
		// returning validation message
		g.JSON(http.StatusBadRequest, gin.H{
			// error message needs to be customized
			"error":err.Error(),
			"data": "",
		})
		return
	}
	// appending some data
	post.Status = "Active"
	db.Create(&post)

	g.JSON(http.StatusCreated, gin.H{
		"error":"",
		"data": post,
	})
}

func updatePost(g *gin.Context) {
	oldPost := getPostById(g)
	if oldPost.ID == 0 {
		return
	}
	var requestPost Post
	// if there is an error; the sent data doesnt match with the validations
	if err := g.ShouldBindJSON(&requestPost); err != nil {
		// returning validation message
		g.JSON(http.StatusBadRequest, gin.H{
			// error message needs to be customized
			"error":err.Error(),
			"data": "",
		})
		return
	}

	oldPost.Title = requestPost.Title
	oldPost.Description = requestPost.Description
	if requestPost.Status != "" {
		oldPost.Status = requestPost.Status
	}
	db.Save(&oldPost)

	g.JSON(http.StatusOK, gin.H{
		"error":"",
		"data": oldPost,
		"message": "updated successfully",
	})
}

func deletePost(g *gin.Context) {
	post := getPostById(g)
	if post.ID == 0 {
		return
	}
	db.Delete(&post, post.ID)
	// for hard delete
	//db.Unscoped().Delete(&post, postId)
	g.JSON(http.StatusOK, gin.H{
		"error":"",
		"data": "",
		"message": "deleted successfully",
	})
}

func getPostById(g *gin.Context)  Post {
	postId := g.Param("id")
	var post Post
	// get the post
	db.First(&post, postId)

	// checks if the post exists in the DB
	// when there is no matched record in DB, GO returns the model (Post here)
	// with the default data for every column type, ex: string => "", integer => 0
	if post.ID == 0 {
		g.JSON(http.StatusNotFound, gin.H{
			"error":"didnt find this post",
			"data": "",
		})
	}
	return post
}