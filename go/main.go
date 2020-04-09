package main

import (
	"fmt"
	"math/rand"
	"time"
)

const matrixLen = 8

func main() {
	fakesum := 0
	iterations := 100_000
	rand.Seed(time.Now().UnixNano())

	// SEQUENTIAL
	var sequentialElapsedTotal time.Duration = 0

	for i := 0; i < iterations; i++ {
		matrix1 := generateMatrix()
		matrix2 := generateMatrix()

		start := time.Now()
		resultingMatrix := matrix(matrix1, matrix2)
		elapsed := time.Since(start)
		sequentialElapsedTotal += elapsed

		fakesum += len(resultingMatrix)
	}

	sequentialMean := time.Duration(int64(sequentialElapsedTotal) / int64(iterations))

	// CONCURRENT
	var concurrentElapsedTotal time.Duration = 0

	for i := 0; i < iterations; i++ {
		matrix1 := generateMatrix()
		matrix2 := generateMatrix()

		start := time.Now()
		resultingMatrix := concurrentMatrix(matrix1, matrix2)
		elapsed := time.Since(start)
		concurrentElapsedTotal += elapsed

		fakesum += len(resultingMatrix)
	}

	concurrentMean := time.Duration(int64(concurrentElapsedTotal) / int64(iterations))

	fmt.Printf("Fakesum: %d\n", fakesum)
	fmt.Printf("Matrix size: %d\n", matrixLen)
	fmt.Printf("Iterations: %d\n\n", iterations)

	fmt.Println("---SEQUENTIAL---")
	fmt.Printf("Total elapsed time: %dns\n", sequentialElapsedTotal.Nanoseconds())
	fmt.Printf("Mean time: %dns\n", sequentialMean.Nanoseconds())

	fmt.Println("---CONCURRENT---")
	fmt.Printf("Total elapsed time: %dns\n", concurrentElapsedTotal.Nanoseconds())
	fmt.Printf("Mean time: %dns\n", concurrentMean.Nanoseconds())
}

func generateMatrix() [][]int {
	var m [][]int = make([][]int, matrixLen)

	for row := range m {
		m[row] = make([]int, matrixLen)
		for col := range m[row] {
			m[row][col] = rand.Intn(100)
		}
	}

	return m
}
