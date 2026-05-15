package services

import (
	"crypto/aes"
	"crypto/cipher"
	"crypto/rand"
	"encoding/hex"
	"io"
	"os"
)

// 32-byte AES key
var secretKey = []byte("12345678901234567890123456789012")

func EncryptFile(inputPath string, outputPath string) error {

	// Open original file
	inputFile, err := os.Open(inputPath)

	if err != nil {
		return err
	}

	defer inputFile.Close()

	// Create encrypted output file
	outputFile, err := os.Create(outputPath)

	if err != nil {
		return err
	}

	defer outputFile.Close()

	// Create AES cipher block
	block, err := aes.NewCipher(secretKey)

	if err != nil {
		return err
	}

	// Create random IV
	iv := make([]byte, aes.BlockSize)

	_, err = io.ReadFull(rand.Reader, iv)

	if err != nil {
		return err
	}

	// Save IV at beginning of encrypted file
	_, err = outputFile.Write(iv)

	if err != nil {
		return err
	}

	// Create encryption stream
	stream := cipher.NewCFBEncrypter(block, iv)

	// Stream writer
	writer := &cipher.StreamWriter{
		S: stream,
		W: outputFile,
	}

	// Encrypt file while streaming
	_, err = io.Copy(writer, inputFile)

	if err != nil {
		return err
	}

	return nil
}

// Optional helper to view key
func GetEncryptionKey() string {
	return hex.EncodeToString(secretKey)
}

func DecryptFile(inputPath string, outputPath string) error {

	// Open encrypted file
	inputFile, err := os.Open(inputPath)

	if err != nil {
		return err
	}

	defer inputFile.Close()

	// Create output decrypted file
	outputFile, err := os.Create(outputPath)

	if err != nil {
		return err
	}

	defer outputFile.Close()

	// Create AES cipher
	block, err := aes.NewCipher(secretKey)

	if err != nil {
		return err
	}

	// Read IV
	iv := make([]byte, aes.BlockSize)

	_, err = io.ReadFull(inputFile, iv)

	if err != nil {
		return err
	}

	// Create decryption stream
	stream := cipher.NewCFBDecrypter(block, iv)

	// Stream reader
	reader := &cipher.StreamReader{
		S: stream,
		R: inputFile,
	}

	// Decrypt while copying
	_, err = io.Copy(outputFile, reader)

	if err != nil {
		return err
	}

	return nil
}
