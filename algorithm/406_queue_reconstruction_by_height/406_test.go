package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestAllPathsSourceTarget(t *testing.T) {
	assert := assert.New(t)

	dataset := [][][]int{
		[][]int{{7, 0}, {4, 4}, {7, 1}, {5, 0}, {6, 1}, {5, 2}},
	}

	expected := [][][]int{
		[][]int{{5, 0}, {7, 0}, {5, 2}, {6, 1}, {4, 4}, {7, 1}},
	}

	for index := range dataset {
		path := reconstructQueue(dataset[index])
		assert.Equal(path, expected[index], "not correct.")
	}
}
