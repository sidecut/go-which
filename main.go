package main

import (
	"fmt"
	"os"
	"strings"
)

func main() {
	args := os.Args
	path := os.Getenv("PATH")
	pathDirs := strings.Split(path, ":")
	// de-dup
	pathDirs = deduplicate(pathDirs)
	success := true
	for _, arg := range args[1:] {
		if !findFile(pathDirs, arg) {
			fmt.Fprintln(os.Stderr, arg, "not found")
			success = false
		}
	}

	if !success {
		os.Exit(1)
	}
}
