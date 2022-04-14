package main

import (
	"os"
	"strings"
)

func main() {
	args := os.Args
	path := os.Getenv("PATH")
	pathDirs := strings.Split(path, ":")
	// de-dup
	pathDirs = deduplicate(pathDirs)
	for _, arg := range args[1:] {
		findFile(pathDirs, arg)
	}
}
