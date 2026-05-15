package services

import (
	"crypto/sha256"
	"encoding/hex"
	"io"
	"os"
)

func GenerateFileHash(filePath string) (string, error) {

	file, err := os.Open(filePath)

	if err != nil {
		return "", err
	}

	defer file.Close()

	hasher := sha256.New()

	_, err = io.Copy(hasher, file)

	if err != nil {
		return "", err
	}

	hashBytes := hasher.Sum(nil)

	hashString := hex.EncodeToString(hashBytes)

	return hashString, nil
}
