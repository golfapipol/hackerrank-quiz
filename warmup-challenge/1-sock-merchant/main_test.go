package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestSockMerchant(t *testing.T) {
	t.Run("Case 1: Input n = 7, ar = [1,2,1,2,1,3,2] Output 2", func(t *testing.T) {
		expectedResult := int32(2)
		n := int32(7)
		array := []int32{1, 2, 1, 2, 1, 3, 2}

		assert.Equal(t, expectedResult, sockMerchant(n, array))
	})
	t.Run("Case 1: Input n = 9, ar = [10,20,20,10,10,30,50,10,20] Output 3", func(t *testing.T) {
		expectedResult := int32(3)
		n := int32(7)
		array := []int32{10, 20, 20, 10, 10, 30, 50, 10, 20}

		assert.Equal(t, expectedResult, sockMerchant(n, array))
	})
}
