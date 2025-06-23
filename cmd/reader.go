package cmd

import (
	"fmt"
	"os"
)

func ReadFileContent(filepath string) []byte {
	// check if file exists
	contents, err := os.ReadFile(filepath)
	if err != nil {
		fmt.Println("file does not exists at specified path")
		os.Exit(1)
	}
	return contents
}
