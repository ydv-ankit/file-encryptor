package main

import (
	"fmt"

	"github.com/ydv-ankit/file-encryptor/cmd"
)

func main() {
	// get args from cli
	args := cmd.ExtractArgs()
	fmt.Println(args)
	var fileContent []byte
	if args.Filepath != "" {
		fileContent = cmd.ReadFileContent(args.Filepath)
		fmt.Println(string(fileContent))
	}
}
