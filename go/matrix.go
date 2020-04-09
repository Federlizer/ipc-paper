package main

func matrix(matrix1, matrix2 [][]int) [][]int {
	var result [][]int = make([][]int, matrixLen)

	for col := range matrix1 {
		result[col] = make([]int, matrixLen)
		for row := range result[col] {
			result[col][row] = matrix1[col][row] * matrix2[col][row]
		}
	}

	return result
}
