package main

func multiplyArray(arr []int) int {
	var total = 1

	for _, val := range arr {
		total *= val
	}

	return total
}
