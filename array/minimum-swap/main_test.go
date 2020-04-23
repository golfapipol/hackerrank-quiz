package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestMinimumSwaps(t *testing.T) {
	t.Run("Case1: 4 3 1 2 Output 3", func(t *testing.T) {
		expectedResult := int32(3)
		list := []int32{4, 3, 1, 2}
		assert.Equal(t, expectedResult, minimumSwaps(list))
	})

	t.Run("Case1: 2 3 4 1 5 Output 3", func(t *testing.T) {
		expectedResult := int32(3)
		list := []int32{2, 3, 4, 1, 5}
		assert.Equal(t, expectedResult, minimumSwaps(list))
	})

	t.Run("Case1: 1 3 5 2 4 6 7 Output 3", func(t *testing.T) {
		expectedResult := int32(3)
		list := []int32{1, 3, 5, 2, 4, 6, 7}
		assert.Equal(t, expectedResult, minimumSwaps(list))
	})
}
