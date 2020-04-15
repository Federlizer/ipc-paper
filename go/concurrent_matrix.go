package main

import (
	"runtime"
)

func concurrentMatrix(matrix1, matrix2 [][]int, matrixLen int) [][]int {
	numCPUs := runtime.NumCPU()
	routineMatrixLength := matrixLen / numCPUs

	var ch chan bool = make(chan bool)
	var result [][]int = make([][]int, matrixLen)

	for row := 0; row < matrixLen; row += routineMatrixLength {
		split := make([][]int, routineMatrixLength)

		for i := 0; i < routineMatrixLength; i++ {
			split[i] = matrix1[row+i]
		}

		go func(row int, ch chan<- bool) {
			for r := range split {
				result[row+r] = make([]int, matrixLen)
				for c := range split[r] {
					result[row+r][c] = split[r][c] * matrix2[row+r][c]
				}
			}
			ch <- true
		}(row, ch)
	}

	for i := 0; i < numCPUs; i++ {
		<-ch
	}

	return result
}
