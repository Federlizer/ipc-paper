package main

func concurrentMatrix(matrix1, matrix2 [][]int) [][]int {
	var ch chan bool = make(chan bool)
	var result [][]int = make([][]int, matrixLen)

	for row := 0; row < matrixLen; row += 2 {
		split := make([][]int, 2)
		split[0], split[1] = matrix1[row], matrix1[row+1]

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

	for i := 0; i < matrixLen/2; i++ {
		<-ch
	}

	return result
}
