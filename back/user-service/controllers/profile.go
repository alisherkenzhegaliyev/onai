package controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func RootRoute(c *gin.Context) {
	c.IndentedJSON(http.StatusOK, gin.H{"message": "Hello wws!"})
}
