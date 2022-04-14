package main

import (
	"os"
	"strings"
)

func main() {
	args := os.Args
	path := os.Getenv("PATH")
	pathDirs := strings.Split(path, ":")
	for _, arg := range args[1:] {
		findFile(pathDirs, arg)
	}
}
