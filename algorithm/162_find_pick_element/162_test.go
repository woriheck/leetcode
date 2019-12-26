package peak

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestFindPeakElement(t *testing.T) {
	assert := assert.New(t)

	dataset := [][]int{
		[]int{1, 2, 3, 1},
		[]int{1, 2, 1, 3, 5, 6, 4},
		[]int{2, 1},
	}

	expected := []int{2, 5, 0}

	for index := range dataset {
		path := findPeakElement(dataset[index])
		assert.Equal(path, expected[index], "Not correct.")
	}
}
