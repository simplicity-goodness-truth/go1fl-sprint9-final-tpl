package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

type GeneratedRandomElements struct {
	size int
	wait int
}

type MaximumElements struct {
	data []int
	wait int
}

func TestGenerateRandomElements(t *testing.T) {

	// Filling structure with size parameter and awaited value
	requests := []GeneratedRandomElements{
		{0, 0},
		{-123, 0},
		{200, 200},
		{11111, 11111},
		{1, 1},
		{4, 4},
		{-2147483648, 0},
		{214748364, 214748364},
		{10,10},
	}

	for _, v := range requests {

		elements := generateRandomElements(v.size)
		assert.Equal(t, v.wait, len(elements))
	}
}

func TestMaximum(t *testing.T) {

	// Filling structure with size parameter and awaited value
	requests := []MaximumElements{
		{data: []int{0, 0}, wait: 0},
		{data: []int{0}, wait: 0},
		{data: []int{}, wait: 0},
		{data: []int{1}, wait: 1},
		{data: []int{454545}, wait: 454545},
		{data: []int{2, 3, 4, 5, 454545}, wait: 454545},
		{data: []int{0, 3, 4, 5, 0}, wait: 5},
		{data: []int{-1}, wait: -1},
		{data: []int{-454545, 0}, wait: 0},
		{data: []int{2, 3, 4, -5, 454545}, wait: 454545},
		{data: []int{0, -3, 4, 5, 0}, wait: 5},
		{data: []int{-1, -3, -4, -5, -120}, wait: -1},
	}

	for _, v := range requests {

		maxElement := maximum(v.data)

		assert.Equal(t, v.wait, maxElement)
	}

}

func TestMaxChunks(t *testing.T) {

	// Filling structure with size parameter and awaited value
	requests := []MaximumElements{
		{data: []int{2, 3, 4, 15, 3, 2, 4, 1, 6, 4, 3, 2, 3, 2, 4, 1, 3, 2, 433}, wait: 433},
		{data: []int{0, 0}, wait: 0},
		{data: []int{}, wait: 0},
		{data: []int{1}, wait: 1},
		{data: []int{454545}, wait: 454545},
		{data: []int{2, 3, 4, 5, 454545}, wait: 454545},
		{data: []int{0, 3, 4, 5, 0}, wait: 5},
		{data: []int{2, 3, 4, 5, 3, 2, 4, 1}, wait: 5},
		{data: []int{2, 3, 4, 5, 3, 2, 4, 1, 0, 0}, wait: 5},
		{data: []int{2, 3, 4, 15, 3, 2, 4, 1, 3, 2}, wait: 15},
		{data: []int{2, 3, 4, 15, 3, 2, 4, 1, 3, 2, 3, 2, 4, 1, 3, 2}, wait: 15},
		{data: []int{2, 3, 4, 15, 3, 2, 4, 1, 3, 2, 3, 2, 4, 1, 3, 222, 43}, wait: 222},
	}

	for _, v := range requests {

		maxElement := maxChunks(v.data)

		assert.Equal(t, v.wait, maxElement)
	}

}
