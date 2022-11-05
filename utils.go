package main

import (
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
)

func wrappContext(ctx *appContext, f func(*gin.Context, *appContext)) func(*gin.Context) {
	return func(c *gin.Context) { f(c, ctx) }
}

func returnRequestError(c *gin.Context, message string, err error) {
	logError(message, err)
	c.IndentedJSON(http.StatusBadRequest, gin.H{"error": message})
}

func returnInternalError(c *gin.Context, message string, err error) {
	logError(message, err)
	c.IndentedJSON(http.StatusInternalServerError, gin.H{"error": message})
}

func logError(message string, err error) {
	if err != nil {
		log.Printf("%s: %v", message, err)
	} else {
		log.Printf("%s", message)
	}
}
