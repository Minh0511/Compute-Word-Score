package main

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestComputeWordScore(t *testing.T) {
	expected := computeWordScore(11, 7, 10)
	assert.Equal(t, "Invalid input", expected)
}
