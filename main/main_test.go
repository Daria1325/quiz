package main

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestUnifyTest(t *testing.T) {
	input := "\nTeSting This FUNcTiOn\n\n"
	expResult := "testing this function"

	actualResult := unifyText(input)
	assert.Equal(t, expResult, actualResult)
}
func TestReadFile(t *testing.T) {
	input := "some.json"

	err := readFile(input)
	if err == nil {
		t.Error("something went wrong, there is has to be an error")
	}
}
func TestSuffledArray(t *testing.T) {
	input := []Question{{"1", "1"}, {"2", "2"}, {"3", "3"}, {"4", "4"}}
	result := []Question{{"1", "1"}, {"2", "2"}, {"3", "3"}, {"4", "4"}}
	shuffleArray(result)
	flag := false
	for i, exepl := range result {
		if exepl != input[i] {
			flag = true
		}
	}
	if !flag {
		t.Error("Array wasn't shuffled")
	}
}
