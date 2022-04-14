package main

import "fmt"

func deduplicate[T comparable](items []T) []T {
	usedMap := make(map[T]int)
	copy := []T{}

	// Record first use of each item
	for i, item := range items {
		if _, used := usedMap[item]; !used {
			usedMap[item] = i
			copy = append(copy, item)
		}
	}

	// TEMP
	fmt.Printf("%v\n", usedMap)

	return copy
}
