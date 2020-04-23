package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestDiagonalDifference(t *testing.T) {
	t.Run(`Case 1`, func(t *testing.T) {
		expectedResult := int32(15)
		input := [][]int32{
			{11, 2, 4},
			{4, 5, 6},
			{10, 8, -12},
		}

		assert.Equal(t, expectedResult, diagonalDifference(input))
	})
}
