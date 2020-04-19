package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestJumpingOnClouds(t *testing.T) {
	t.Run("Case1: Input 0010010 Output 4", func(t *testing.T) {
		expectedResult := int32(4)
		track := []int32{0, 0, 1, 0, 0, 1, 0}

		assert.Equal(t, expectedResult, jumpingOnClouds(track))
	})
	t.Run("Case2: Input 000010 Output 3", func(t *testing.T) {
		expectedResult := int32(3)
		track := []int32{0, 0, 0, 0, 1, 0}

		assert.Equal(t, expectedResult, jumpingOnClouds(track))
	})
	t.Run("Case3: Input 000100 Output 3", func(t *testing.T) {
		expectedResult := int32(3)
		track := []int32{0, 0, 0, 1, 0, 0}

		assert.Equal(t, expectedResult, jumpingOnClouds(track))
	})
}
