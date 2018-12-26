package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestSingleNonDuplicate(t *testing.T) {
	assert := assert.New(t)

	dataset := [][]int{
		[]int{1, 1, 2},
		[]int{1, 1, 2, 3, 3, 4, 4, 8, 8},
		[]int{3, 3, 7, 7, 10, 11, 11},
		[]int{7},
	}

	expected := []int{
		2, 2, 10, 7,
	}

	for index := range dataset {
		path := singleNonDuplicate(dataset[index])
		assert.Equal(path, expected[index], "Not correct.")
	}
}
