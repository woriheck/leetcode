package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestAllPathsSourceTarget(t *testing.T) {
	assert := assert.New(t)

	dataset := [][][]int{
		[][]int{{1, 2}, {3}, {3}, {}},
		[][]int{{3, 1}, {4, 6, 7, 2, 5}, {4, 6, 3}, {6, 4}, {7, 6, 5}, {6}, {7}, {}},
		[][]int{{1, 2}, {3, 4}, {5}, {6}, {6}, {4, 6}, {}},
	}

	expected := [][][]int{
		[][]int{{0, 1, 3}, {0, 2, 3}},
		[][]int{{0, 3, 6, 7}, {0, 3, 4, 7}, {0, 3, 4, 6, 7}, {0, 3, 4, 5, 6, 7}, {0, 1, 4, 7}, {0, 1, 4, 6, 7}, {0, 1, 4, 5, 6, 7}, {0, 1, 6, 7}, {0, 1, 7}, {0, 1, 2, 4, 7}, {0, 1, 2, 4, 6, 7}, {0, 1, 2, 4, 5, 6, 7}, {0, 1, 2, 6, 7}, {0, 1, 2, 3, 6, 7}, {0, 1, 2, 3, 4, 7}, {0, 1, 2, 3, 4, 6, 7}, {0, 1, 2, 3, 4, 5, 6, 7}, {0, 1, 5, 6, 7}},
		[][]int{{0, 1, 3, 6}, {0, 1, 4, 6}, {0, 2, 5, 4, 6}, {0, 2, 5, 6}},
	}

	for index := range dataset {
		path := allPathsSourceTarget(dataset[index])
		assert.Equal(path, expected[index], "Not correct.")
	}
}
