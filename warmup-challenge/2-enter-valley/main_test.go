package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestCountValley(t *testing.T) {
	t.Run("Case1: Input 8,UDDDUDUU Output 1", func(t *testing.T) {
		expectedResult := int32(1)
		number := int32(8)
		track := "UDDDUDUU"

		assert.Equal(t, expectedResult, countingValleys(number, track))
	})
	t.Run("Case2: Input 12,DDUUDDUDUUUD Output 2", func(t *testing.T) {
		expectedResult := int32(1)
		number := int32(8)
		track := "UDDDUDUU"

		assert.Equal(t, expectedResult, countingValleys(number, track))
	})
}
