package handlers

import "github.com/gin-gonic/gin"

func Index(context *gin.Context) {
	// send json response
	context.JSON(200, gin.H{"message": "Hello World!"})
}

func Page(context *gin.Context) {
	page := context.Params.ByName("page")
	context.JSON(200, gin.H{"title": page})
}
