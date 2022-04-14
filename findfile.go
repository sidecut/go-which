package main

import (
	"fmt"
	"os"
	"path"
)

func findFile(pathDirs []string, filename string) {
	for _, pathDir := range pathDirs {
		fullPath := path.Join(pathDir, filename)
		fileInfo, err := os.Stat(fullPath)
		if err != nil {
			if err.(*os.PathError) == nil {
				// it's not a bad path
				fmt.Fprintln(os.Stderr, err.Error())
			}
		} else {
			if mode := fileInfo.Mode(); mode.IsRegular() {
				if mode&0111 != 0 {
					fmt.Println(fullPath)
				}
			}
		}
	}
}
