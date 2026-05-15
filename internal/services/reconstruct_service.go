package services

import (
	"io"
	"os"
)

func ReconstructFile(chunkPaths []string, outputPath string) error {

	// Create reconstructed output file
	outputFile, err := os.Create(outputPath)

	if err != nil {
		return err
	}

	defer outputFile.Close()

	// Read chunks sequentially
	for _, chunkPath := range chunkPaths {

		chunkFile, err := os.Open(chunkPath)

		if err != nil {
			return err
		}

		// Copy chunk bytes into output file
		_, err = io.Copy(outputFile, chunkFile)

		chunkFile.Close()

		if err != nil {
			return err
		}
	}

	return nil
}
