package main

import (
	"io/ioutil"
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestCheckMagazine(t *testing.T) {
	t.Run("Case 1:", func(t *testing.T) {
		expectedResult := "Yes"
		magazine := []string{"give", "me", "one", "grand", "today", "night"}
		note := []string{"give", "one", "grand", "today"}
		assert.Equal(t, expectedResult, checkMagazine(magazine, note))
	})
	t.Run("Case 2:", func(t *testing.T) {
		expectedResult := "No"
		magazine := []string{"two", "times", "three", "is", "not", "four"}
		note := []string{"two", "times", "two", "is", "four"}

		assert.Equal(t, expectedResult, checkMagazine(magazine, note))
	})
	t.Run("Case 2:", func(t *testing.T) {
		expectedResult := "No"
		magazine := []string{"two", "times", "three", "is", "not", "four"}
		note := []string{"two", "times", "two", "is", "four"}

		assert.Equal(t, expectedResult, checkMagazine(magazine, note))
	})

	t.Run("Case 3:", func(t *testing.T) {
		expectedResult := "No"
		magazine := []string{"two", "times", "three", "is", "not", "four"}
		note := []string{"two", "times", "two", "is", "four"}

		assert.Equal(t, expectedResult, checkMagazine(magazine, note))
	})

	t.Run("Case 13:", func(t *testing.T) {
		expectedResult := "Yes"
		magData, _ := ioutil.ReadFile("magtc13.txt")
		noteData, _ := ioutil.ReadFile("notetc13.txt")
		magazine := strings.Split(string(magData), " ")
		note := strings.Split(string(noteData), " ")

		assert.Equal(t, expectedResult, checkMagazine(magazine, note))
	})
}
