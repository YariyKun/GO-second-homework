package main

import (
	"fmt"
	"math/rand"
	"sync"
	"time"
)

const (
	SIZE   = 100_000_000
	CHUNKS = 8
)

// generateRandomElements
func generateRandomElements(size int) []int {
	if size <= 0 {
		return []int{}
	}

	slice := make([]int, size)
	for i := 0; i < size; i++ {
		slice[i] = rand.Intn(1000000) + 1
	}
	return slice
}

// maximum возвращает максимальное число в слайсе.
func maximum(data []int) int {
	// Обработка крайнего случая: пустой слайс
	if len(data) == 0 {
		return 0
	}

	maxVal := data[0]
	for _, v := range data {
		if v > maxVal {
			maxVal = v
		}
	}
	return maxVal
}

// maxChunks
func maxChunks(data []int) int {
	if len(data) == 0 {
		return 0
	}

	if len(data) < CHUNKS {
		return maximum(data)
	}

	chunkSize := len(data) / CHUNKS
	maxValues := make([]int, CHUNKS)
	var wg sync.WaitGroup

	for i := 0; i < CHUNKS; i++ {
		wg.Add(1)

		startIndex := i * chunkSize
		endIndex := startIndex + chunkSize

		if i == CHUNKS-1 {
			endIndex = len(data)
		}

		go func(index int, chunk []int) {
			defer wg.Done()
			maxValues[index] = maximum(chunk)
		}(i, data[startIndex:endIndex])
	}

	wg.Wait()

	return maximum(maxValues)
}

func main() {
	fmt.Printf("Генерируем %d целых чисел\n", SIZE)
	data := generateRandomElements(SIZE)

	fmt.Println("Ищем максимальное значение в один поток")
	startSingle := time.Now()
	max := maximum(data)
	elapsed := time.Since(startSingle).Microseconds()

	fmt.Printf("Максимальное значение элемента: %d\nВремя поиска: %d мкс\n", max, elapsed)

	fmt.Printf("Ищем максимальное значение в %d потоков\n", CHUNKS)
	startMulti := time.Now()
	max = maxChunks(data)
	elapsed = time.Since(startMulti).Microseconds()

	fmt.Printf("Максимальное значение элемента: %d\nВремя поиска: %d мкс\n", max, elapsed)
}
