package main

import "testing"

func TestIsValid(t *testing.T) {
	dataset := []string{
		"()",
		"{}",
		"[]",
		"(){}[]",
		"{[()]}",
	}

	for _, data := range dataset {
		valid := isValid(data)
		if valid == false {
			t.Errorf("%s was incorrect, expected: true, result: false.", data)
		}
	}
}
