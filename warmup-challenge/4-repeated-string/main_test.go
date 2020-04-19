package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestRepeatedString(t *testing.T) {
	t.Run("Case1: Input aba,10 Output 7", func(t *testing.T) {
		expectedResult := int64(7)
		strings := "aba"
		repeat := int64(10)

		assert.Equal(t, expectedResult, repeatedString(strings, repeat))
	})
	t.Run("Case2: Input a,1000000000000 Output 1000000000000", func(t *testing.T) {
		expectedResult := int64(1000000000000)
		strings := "a"
		repeat := int64(1000000000000)

		assert.Equal(t, expectedResult, repeatedString(strings, repeat))
	})
	t.Run("Case3: Input 000100 Output 3", func(t *testing.T) {
		expectedResult := int64(4)
		strings := "abcac"
		repeat := int64(10)

		assert.Equal(t, expectedResult, repeatedString(strings, repeat))
	})
}
