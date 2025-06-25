package main

import (
	"crypto/sha256"
	"fmt"

	"github.com/ydv-ankit/file-encryptor/cmd"
	"github.com/ydv-ankit/file-encryptor/tea"
)

func main() {
	// get args from cli
	args := cmd.ExtractArgs()

	// Read file content
	fileContent := cmd.ReadFileContent(args.Filepath)
	fmt.Printf("Processing file: %s (%d bytes)\n", args.Filepath, len(fileContent))

	// Read key file
	key := cmd.ReadFileContent(args.KeyfilePath)

	// convert key to 128 bit hash
	h := sha256.New()
	h.Write(key)
	keyHash := h.Sum(nil)[:16]

	if args.Encrypt {
		cipherText := tea.EncryptData(fileContent, keyHash)
		cmd.WriteFileContent(args.Filepath, cipherText)
	}
	if args.Decrypt {
		plainText := tea.DecryptData(fileContent, keyHash)
		cmd.WriteFileContent(args.Filepath, plainText)
	}
	fmt.Println("operation successful")
	fmt.Println("data written to file:", args.Filepath)
}
