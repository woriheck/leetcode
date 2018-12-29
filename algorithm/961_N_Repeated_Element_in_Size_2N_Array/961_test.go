package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestRepeatedNTimes(t *testing.T) {
	assert := assert.New(t)

	dataset := [][]int{
		[]int{1, 2, 3, 3},
		[]int{2, 1, 2, 5, 3, 2},
		[]int{5, 1, 5, 2, 5, 3, 5, 4},
	}

	expected := []int{3, 2, 5}

	for index := range dataset {
		path := repeatedNTimes(dataset[index])
		assert.Equal(path, expected[index], "Not correct.")
	}
}
