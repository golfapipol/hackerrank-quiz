package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestHourglassSum(t *testing.T) {
	t.Run("Case1: Output 28", func(t *testing.T) {
		expectedResult := int32(28)
		matrix := [][]int32{
			{-9, -9, -9, 1, 1, 1},
			{0, -9, 0, 4, 3, 2},
			{-9, -9, -9, 1, 2, 3},
			{0, 0, 8, 6, 6, 0},
			{0, 0, 0, -2, 0, 0},
			{0, 0, 1, 2, 4, 0},
		}
		assert.Equal(t, expectedResult, hourglassSum(matrix))
	})

	t.Run("Case2: Output 19", func(t *testing.T) {
		expectedResult := int32(19)
		matrix := [][]int32{
			{1, 1, 1, 0, 0, 0},
			{0, 1, 0, 0, 0, 0},
			{1, 1, 1, 0, 0, 0},
			{0, 0, 2, 4, 4, 0},
			{0, 0, 0, 2, 0, 0},
			{0, 0, 1, 2, 4, 0},
		}
		assert.Equal(t, expectedResult, hourglassSum(matrix))
	})
	t.Run("Case3: Output -19", func(t *testing.T) {
		expectedResult := int32(-19)
		matrix := [][]int32{
			{0, -4, -6, 0, -7, -6},
			{-1, -2, -6, -8, -3, -1},
			{-8, -4, -2, -8, -8, -6},
			{-3, -1, -2, -5, -7, -4},
			{-3, -5, -3, -6, -6, -6},
			{-3, -6, 0, -8, -6, -7},
		}
		assert.Equal(t, expectedResult, hourglassSum(matrix))
	})
}
