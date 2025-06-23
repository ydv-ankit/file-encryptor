package cmd

import (
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
		case "--encrypt":
			extractedArgs.Encrypt = true
			// toggle decrypt if encrypt is present
			extractedArgs.Decrypt = false
		case "--decrypt":
			extractedArgs.Decrypt = true
			// toggle encrypt if decrypt is present
			extractedArgs.Encrypt = false
		}
		// for key file & main file to process
		if len(arg) > 6 {
			if arg[0:6] == "file=" {
				extractedArgs.Filepath = arg[6:]
			}
			if arg[0:4] == "key=" {
				extractedArgs.KeyfilePath = arg[4:]
			}
		}
	}

	return extractedArgs
}
