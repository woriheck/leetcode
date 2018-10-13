package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestAllPathsSourceTarget(t *testing.T) {
	assert := assert.New(t)

	dataset := [][]string{
		[]string{"kqep", "pekeq"},
	}

	expected := []string{
		"kqeep",
	}

	for index := range dataset {
		path := customSortString(dataset[index][0], dataset[index][1])
		assert.Equal(path, expected[index], "not correct.")
	}
}
