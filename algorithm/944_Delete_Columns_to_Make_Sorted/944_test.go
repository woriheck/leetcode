package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestMinDeletionSize(t *testing.T) {
	assert := assert.New(t)

	dataset := [][]string{
		{"cba", "daf", "ghi"},
		{"a", "b"},
		{"zyx", "wvu", "tsr"},
	}

	expected := []int{
		1, 0, 3,
	}

	for index := range dataset {
		path := minDeletionSize(dataset[index])
		assert.Equal(path, expected[index], "Not correct.")
	}
}
