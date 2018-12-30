package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestFlipAndInvertImage(t *testing.T) {
	assert := assert.New(t)

	dataset := [][][]int{
		[][]int{{1, 1, 0}, {1, 0, 1}, {0, 0, 0}},
		[][]int{{1, 1, 0, 0}, {1, 0, 0, 1}, {0, 1, 1, 1}, {1, 0, 1, 0}},
	}

	expected := [][][]int{
		[][]int{{1, 0, 0}, {0, 1, 0}, {1, 1, 1}},
		[][]int{{1, 1, 0, 0}, {0, 1, 1, 0}, {0, 0, 0, 1}, {1, 0, 1, 0}},
	}

	for index := range dataset {
		data := flipAndInvertImage(dataset[index])
		assert.Equal(data, expected[index], "Not correct.")
	}
}
