package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestDedup(t *testing.T) {
	// rand.Seed(time.Now().UnixNano())
	// for testRun := 0; testRun < 100; testRun++ {
	// 	arrayLength := rand.Intn(100)
	// 	array := make([]int, arrayLength)
	// 	for i := range array {

	// 	}
	// }

	arraySimple := []string{"A", "B", "C", "D", "E"}
	arrayDeduplicated := deduplicate(arraySimple)
	assert.Equal(t, 5, len(arrayDeduplicated))
	assert.EqualValues(t, arraySimple, arrayDeduplicated)

	arraySkipFront := []string{"A", "A", "B", "C", "D", "E"}
	arrayDeduplicated = deduplicate(arraySkipFront)
	assert.Equal(t, 5, len(arrayDeduplicated))
	assert.EqualValues(t, arraySimple, arrayDeduplicated)

	arraySkipMiddle := []string{"A", "B", "C", "C", "D", "E"}
	arrayDeduplicated = deduplicate(arraySkipMiddle)
	assert.Equal(t, 5, len(arrayDeduplicated))
	assert.EqualValues(t, arraySimple, arrayDeduplicated)

	arraySkipEnd := []string{"A", "B", "C", "D", "E", "E", "E"}
	arrayDeduplicated = deduplicate(arraySkipEnd)
	assert.Equal(t, 5, len(arrayDeduplicated))
	assert.EqualValues(t, arraySimple, arrayDeduplicated)

	arrayAllDupes := []string{"A", "A", "A", "A", "A", "A", "A"}
	arrayDeduplicated = deduplicate(arrayAllDupes)
	assert.Equal(t, 1, len(arrayDeduplicated))
	assert.EqualValues(t, []string{"A"}, arrayDeduplicated)

	arrayEmpty := []string{}
	arrayDeduplicated = deduplicate(arrayEmpty)
	assert.Equal(t, 0, len(arrayDeduplicated))
	assert.EqualValues(t, []string{}, arrayDeduplicated)
}
