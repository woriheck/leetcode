package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestSpiralMatrixIII(t *testing.T) {
	assert := assert.New(t)

	dataset := [][]int{
		[]int{1, 4, 0, 0},
		[]int{5, 6, 1, 4},
	}

	expected := [][][]int{
		[][]int{{0, 0}, {0, 1}, {0, 2}, {0, 3}},
		[][]int{{1, 4}, {1, 5}, {2, 5}, {2, 4}, {2, 3}, {1, 3}, {0, 3}, {0, 4}, {0, 5}, {3, 5}, {3, 4}, {3, 3}, {3, 2}, {2, 2}, {1, 2}, {0, 2}, {4, 5}, {4, 4}, {4, 3}, {4, 2}, {4, 1}, {3, 1}, {2, 1}, {1, 1}, {0, 1}, {4, 0}, {3, 0}, {2, 0}, {1, 0}, {0, 0}},
	}

	for index := range dataset {
		path := spiralMatrixIII(dataset[index][0], dataset[index][1], dataset[index][2], dataset[index][3])
		assert.Equal(path, expected[index], "Path not correct.")
	}
}
