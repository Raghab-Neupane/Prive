package services

import (
	"fmt"
	"io"
	"os"
	"path/filepath"
)

func SplitFileIntoChunks(filePath string, chunkSize int64) ([]string, error) {

	// Open encrypted file
	file, err := os.Open(filePath)

	if err != nil {
		return nil, err
	}

	defer file.Close()

	// Create chunks directory
	err = os.MkdirAll("chunks", os.ModePerm)

	if err != nil {
		return nil, err
	}

	var chunkPaths []string

	buffer := make([]byte, chunkSize)

	chunkIndex := 0

	for {

		// Read bytes from file
		bytesRead, err := file.Read(buffer)

		if err != nil && err != io.EOF {
			return nil, err
		}

		// Stop if no more bytes
		if bytesRead == 0 {
			break
		}

		// Create chunk filename
		chunkFileName := fmt.Sprintf("chunk_%d.part", chunkIndex)

		chunkPath := filepath.Join("chunks", chunkFileName)

		// Create chunk file
		chunkFile, err := os.Create(chunkPath)

		if err != nil {
			return nil, err
		}

		// Write bytes into chunk file
		_, err = chunkFile.Write(buffer[:bytesRead])

		if err != nil {
			chunkFile.Close()
			return nil, err
		}

		chunkFile.Close()

		chunkPaths = append(chunkPaths, chunkPath)

		chunkIndex++
	}

	return chunkPaths, nil
}
