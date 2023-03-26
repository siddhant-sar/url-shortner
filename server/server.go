package server

import "github.com/gin-gonic/gin"

func InitGinServer() {
	router := gin.Default()
	router.POST("/add", NewUrlObject)
	router.Run()
}
