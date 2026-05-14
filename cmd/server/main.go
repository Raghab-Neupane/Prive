package main

import (
	"Prive/internal/api"

	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()

	r.POST("/upload", api.UploadHandler)
	r.Run(":8080")
}
