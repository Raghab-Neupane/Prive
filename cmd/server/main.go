package main

import (
	"Prive/internal/api"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()

	r.Use(cors.Default())

	r.POST("/upload", api.UploadHandler)

	r.Run(":8080")
}
