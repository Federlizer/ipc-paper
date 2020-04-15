package main

import (
	"fmt"
	"math/rand"
	"runtime"
	"time"
)

func main() {
	rand.Seed(time.Now().UnixNano())
	matrixLen := 8

	for matrixLen <= 2048 {
		fmt.Printf("%-20s %8s %11s\n\n", "name", "mean(ns)", "iterations")
		iterations := 2

		for iterations <= 1048576 {
			fakesum := 0
			matrix1 := generateMatrix(matrixLen)
			matrix2 := generateMatrix(matrixLen)

			start := time.Now()

			for i := 0; i < iterations; i++ {
				fakesum += len(concurrentMatrix(matrix1, matrix2, matrixLen))
			}

			elapsed := time.Since(start)
			meanTime := time.Duration(int64(elapsed) / int64(iterations))

			fmt.Printf("%-20s %8d %11d\n",
				fmt.Sprintf("matrix length %d", matrixLen),
				meanTime,
				iterations,
			)

			iterations *= 2
		}
		fmt.Printf("\n---MATRIX DOUBLED---\n")
		matrixLen *= 2
	}
}

func generateMatrix(matrixLen int) [][]int {
	var m [][]int = make([][]int, matrixLen)

	for row := range m {
		m[row] = make([]int, matrixLen)
		for col := range m[row] {
			m[row][col] = rand.Intn(100)
		}
	}

	return m
}
