package main

import (
	"os"

	"github.com/dooomit/file-converter/cmd"
)

func main() {
	err := cmd.Execute()
	if err != nil {
		os.Exit(1)
	}
}
