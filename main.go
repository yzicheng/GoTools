package main

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func main() {
	router := gin.Default()
	goodsgroup := router.Group("/goods")
	{
		goodsgroup.GET("/welcome", welcome)
	}
	err := router.Run(":8083")
	if err != nil {
		return
	}
}

func formPost(context *gin.Context) {
	message := context.PostForm("message")
	nic := context.DefaultPostForm("nic", "nic")
	context.JSON(http.StatusOK, gin.H{
		"message": message,
		"nic":     nic,
	})
}

func welcome(c *gin.Context) {
	c.JSON(http.StatusOK, "hello,world")
}
