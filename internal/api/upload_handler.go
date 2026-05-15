package api

import (
	"Prive/internal/services"
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

	// Generate SHA-256 hash
	hash, err := services.GenerateFileHash(filePath)

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": "Failed to generate hash",
		})
		return
	}

	encryptedFilePath := filepath.Join("uploads", file.Filename+".encrypted")
	err = services.EncryptFile(filePath, encryptedFilePath)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": "Failed to encrypt file",
		})
		return
	}

	chunkPaths, err := services.SplitFileIntoChunks(encryptedFilePath, 1024*100)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": "Failed to split file into chunks",
		})
		return
	}

	reconstructedPath := "reconstructed/rebuilt.encrypted"

	os.MkdirAll("reconstructed", os.ModePerm)

	err = services.ReconstructFile(chunkPaths, reconstructedPath)

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": "Failed to reconstruct file",
		})
		return
	}

	decryptedPath := filepath.Join(
		"reconstructed",
		"restored_"+file.Filename,
	)

	err = services.DecryptFile(reconstructedPath, decryptedPath)

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": "Failed to decrypt reconstructed file",
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message":           "File uploaded successfully",
		"hash":              hash,
		"encryptedFile":     encryptedFilePath,
		"encryptionKey":     services.GetEncryptionKey(),
		"chunks":            chunkPaths,
		"reconstructedFile": reconstructedPath,
		"decryptedFile":     decryptedPath,
	})
}
