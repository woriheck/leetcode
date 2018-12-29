package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestSortArrayByParity(t *testing.T) {
	assert := assert.New(t)

	dataset := [][]int{
		[]int{3, 1, 2, 4},
	}

	expected := [][]int{
		[]int{4, 2, 3, 1},
	}

	for index := range dataset {
		path := sortArrayByParity(dataset[index])
		assert.Equal(path, expected[index], "Not correct.")
	}
}
