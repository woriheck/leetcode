package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestIsValid(t *testing.T) {
	assert := assert.New(t)

	dataset := []string{
		"()",
		"{}",
		"[]",
		"(){}[]",
		"{[()]}",
	}

	for _, data := range dataset {
		valid := isValid(data)
		assert.True(valid, "Not correct.")
	}
}
