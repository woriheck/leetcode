package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestCanVisitAllRooms(t *testing.T) {
	assert := assert.New(t)

	dataset := [][][]int{
		[][]int{{1}, {2}, {3}, {}},
		[][]int{{1, 3}, {3, 0, 1}, {2}, {0}},
	}

	expected := []bool{
		true,
		false,
	}

	for index := range dataset {
		path := canVisitAllRooms(dataset[index])
		assert.Equal(path, expected[index], "Not correct.")
	}
}
