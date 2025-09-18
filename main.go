package main

import (
	"fmt"
	"math/rand"
	"slices"
	"sync"
	"time"
)

const (
	SIZE   = 100_000_000
	CHUNKS = 8
)

var wg sync.WaitGroup

// generateRandomElements generates random elements.
func generateRandomElements(size int) []int {

	// Declaring a slice
	var elements []int

	// Validation for negative and zero size
	if size == 0 {
		return elements
	}

	// Preparing a random number generator
	rnd := rand.New(rand.NewSource(1001))

	// Looping through amount of elements
	for i := 0; i < size; i++ {

		// Generating a random number
		randomNumber := rnd.Int()

		// Appending a random number to a slice
		elements = append(elements, randomNumber)

	}

	return elements

}

// maximum returns the maximum number of elements.
func maximum(data []int) int {

	// Validation for an empty slice
	if len(data) == 0 {
		return 0
	}

	// Validation for a slice with one element
	if len(data) == 1 {
		return data[0]
	}

	// Searching for a last element
	maxElement := slices.Max(data)

	return maxElement

}

// maxChunks returns the maximum number of elements in a chunks.
func maxChunks(data []int) int {

	// Validation for an empty slice
	if len(data) == 0 {
		return 0
	}

	// Validation for a slice with one element
	if len(data) == 1 {
		return data[0]
	}

	// Validation for zero chunks amount
	if CHUNKS == 0 {
		return 0
	}

	// Slice to store maximum values, identified by dedicated Go routines
	maxValues := make([]int, CHUNKS)

	// Chunk storage
	var chunk []int

	// Calculation of a chunk size
	// If we have less data, than Goroutine can process, then a single Gorotutine can process all data
	chunkSize := len(data)

	if CHUNKS <= chunkSize {
		chunkSize = (len(data) / CHUNKS)
	}

	// Executing Go routines
	for i := 0; i < CHUNKS; i++ {

		wg.Add(1)

		// Calculating chunk's start and end positions
		chunkStartIndex := i * chunkSize
		chunkEndIndex := (chunkStartIndex + chunkSize)

		// If we are executing a last Goroutine, and there is still data to process, add this leftover to a chunk
		// Otherwise compose a chunk using start and end positions

		if (i == CHUNKS-1) && (chunkEndIndex < len(data)) {
			chunk = data[chunkStartIndex:len(data)]
		} else {
			chunk = data[chunkStartIndex:chunkEndIndex]
		}

		// Executing a single Goroutine
		go func(index int, chunk []int) {

			defer wg.Done()

			// Calculating a max value for a chunk and store it in a slice of maximum values
			maxValues[index] = maximum(chunk)

		}(i, chunk)
	}

	wg.Wait()

	// Returing a maximum from a slice of maximum values
	return maximum(maxValues)
}

func main() {
	fmt.Printf("Генерируем %d целых чисел\n", SIZE)

	data := generateRandomElements(SIZE)

	fmt.Println("Ищем максимальное значение в один поток")

	start := time.Now()
	max := maximum(data)
	elapsed := time.Since(start).Microseconds()

	fmt.Printf("Максимальное значение элемента: %d\nВремя поиска: %d ms\n", max, elapsed)

	fmt.Printf("Ищем максимальное значение в %d потоков\n", CHUNKS)

	start = time.Now()
	max = maxChunks(data)
	elapsed = time.Since(start).Microseconds()

	fmt.Printf("Максимальное значение элемента: %d\nВремя поиска: %d ms\n", max, elapsed)
}
