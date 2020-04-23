package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestRotLeft(t *testing.T) {
	t.Run("Case1: 1 2 3 4 5, 4 Output 5,1,2,3,4", func(t *testing.T) {
		expectedResult := []int32{5, 1, 2, 3, 4}
		list := []int32{1, 2, 3, 4, 5}
		assert.Equal(t, expectedResult, rotLeft(list, 4))
	})
}
