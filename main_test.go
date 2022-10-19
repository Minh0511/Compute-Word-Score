package main

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestComputeWordScoreLine2And3True(t *testing.T) {
	expected := computeWordScore(-1, 0, 0)
	assert.Equal(t, "Invalid input", expected)
}

func TestComputeWordScoreLine2And3False(t *testing.T) {
	expected := computeWordScore(1, 1, 1)
	assert.Equal(t, "F", expected)
}

func TestComputeWordScoreLine7False(t *testing.T) {
	expected := computeWordScore(1, 1, 8)
	assert.Equal(t, "F", expected)
}

func TestComputeWordScoreLine12True(t *testing.T) {
	expected := computeWordScore(1, 1, 10)
	assert.Equal(t, "F", expected)
}

func TestComputeWordScoreLine12False(t *testing.T) {
	expected := computeWordScore(9, 9, 9)
	assert.Equal(t, "A+", expected)
}

func TestComputeWordScoreLine14True(t *testing.T) {
	expected := computeWordScore(1, 6, 10)
	assert.Equal(t, "D", expected)
}

func TestComputeWordScoreLine16True(t *testing.T) {
	expected := computeWordScore(2, 6, 10)
	assert.Equal(t, "D+", expected)
}

func TestComputeWordScoreLine18True(t *testing.T) {
	expected := computeWordScore(1, 7, 10)
	assert.Equal(t, "C", expected)
}

func TestComputeWordScoreLine20True(t *testing.T) {
	expected := computeWordScore(1, 9, 10)
	assert.Equal(t, "C+", expected)
}

func TestComputeWordScoreLine22True(t *testing.T) {
	expected := computeWordScore(1, 10, 10)
	assert.Equal(t, "B", expected)
}

func TestComputeWordScoreLine24True(t *testing.T) {
	expected := computeWordScore(4, 10, 10)
	assert.Equal(t, "B+", expected)
}

func TestComputeWordScoreLine26True(t *testing.T) {
	expected := computeWordScore(5, 10, 10)
	assert.Equal(t, "A", expected)
}

func TestComputeWordScoreLine30True(t *testing.T) {
	expected := computeWordScore(9, 9, 9)
	assert.Equal(t, "A+", expected)
}
