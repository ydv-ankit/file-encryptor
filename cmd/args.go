package cmd

import (
	"fmt"
	"os"
)

type Args struct {
	Encrypt     bool
	Decrypt     bool
	Filepath    string
	KeyfilePath string
}

func ExtractArgs() Args {
	args := os.Args
	extractedArgs := Args{}
	for i := 1; i < len(args); i++ {
		arg := args[i]
		// encrypt/decrypt
		switch arg {
		case "-e":
			extractedArgs.Encrypt = true
			// toggle decrypt if encrypt is present
			extractedArgs.Decrypt = false
		case "-d":
			extractedArgs.Decrypt = true
			// toggle encrypt if decrypt is present
			extractedArgs.Encrypt = false
		}
		// for key file & main file to process
		if len(arg) > 5 {
			if arg[0:5] == "file=" {
				extractedArgs.Filepath = arg[5:]
			}
			if arg[0:4] == "key=" {
				extractedArgs.KeyfilePath = arg[4:]
			}
		}
	}

	// validate args
	if extractedArgs.Filepath == "" {
		fmt.Println("Error: No file specified. Use file=<path> to specify input file.")
		os.Exit(1)
	}

	if extractedArgs.KeyfilePath == "" {
		fmt.Println("Error: No key file specified. Use key=<path> to specify input file.")
		os.Exit(1)
	}

	if !extractedArgs.Encrypt && !extractedArgs.Decrypt {
		fmt.Println("Error: Please specify --encrypt or --decrypt")
		os.Exit(1)
	}
	return extractedArgs
}
