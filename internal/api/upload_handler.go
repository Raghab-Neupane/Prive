package api

import (
	"net/http"
	"os"
	"path/filepath"

	"github.com/gin-gonic/gin"
)

func UploadHandler(c *gin.Context) {

	// Get uploaded file
	file, err := c.FormFile("file")

	// If file not received
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "No file uploaded",
		})
		return
	}

	// Create uploads folder if it doesn't exist
	os.MkdirAll("uploads", os.ModePerm)

	// Create destination path
	filePath := filepath.Join("uploads", file.Filename)

	// Save uploaded file
	err = c.SaveUploadedFile(file, filePath)

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": "Failed to save file",
		})
		return
	}

	// Success response
	c.JSON(http.StatusOK, gin.H{
		"message": "File uploaded successfully",
		"file":    file.Filename,
		"path":    filePath,
	})
}
