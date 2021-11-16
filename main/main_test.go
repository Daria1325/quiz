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
