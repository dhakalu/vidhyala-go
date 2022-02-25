package server

import (
	gin "github.com/gin-gonic/gin"
)

func StartApplication() *gin.Engine {
	server := gin.Default()
	return server
}
