package cmd

import (
	"os"
)

func ReadFileContent(filepath string) []byte {
	contents, err := os.ReadFile(filepath)
	if err != nil {
		panic("error reading file")
	}
	return contents
}

func WriteFileContent(filepath string, data []byte) {
	_, err := os.Create(filepath)
	if err != nil {
		panic(err)
	}
	f, err := os.OpenFile(filepath, os.O_WRONLY, os.FileMode(os.O_WRONLY))
	if err != nil {
		panic(err)
	}
	_, err = f.Write(data)
	if err != nil {
		panic(err)
	}
	f.Close()
}
