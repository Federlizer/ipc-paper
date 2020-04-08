package main

import (
	"fmt"
	"math/big"
	"time"
)

var matrix = [][]*big.Int{
	{big.NewInt(1), big.NewInt(2), big.NewInt(3), big.NewInt(4), big.NewInt(5), big.NewInt(6), big.NewInt(7), big.NewInt(8), big.NewInt(9)},
	{big.NewInt(10), big.NewInt(11), big.NewInt(12), big.NewInt(13), big.NewInt(14), big.NewInt(15), big.NewInt(16), big.NewInt(17), big.NewInt(18), big.NewInt(19)},
	{big.NewInt(20), big.NewInt(21), big.NewInt(22), big.NewInt(23), big.NewInt(24), big.NewInt(25), big.NewInt(26), big.NewInt(27), big.NewInt(28), big.NewInt(29)},
	{big.NewInt(30), big.NewInt(31), big.NewInt(32), big.NewInt(33), big.NewInt(34), big.NewInt(35), big.NewInt(36), big.NewInt(37), big.NewInt(38), big.NewInt(39)},
	{big.NewInt(40), big.NewInt(41), big.NewInt(42), big.NewInt(43), big.NewInt(44), big.NewInt(45), big.NewInt(46), big.NewInt(47), big.NewInt(48), big.NewInt(49)},
	{big.NewInt(50), big.NewInt(51), big.NewInt(52), big.NewInt(53), big.NewInt(54), big.NewInt(55), big.NewInt(56), big.NewInt(57), big.NewInt(58), big.NewInt(59)},
	{big.NewInt(60), big.NewInt(61), big.NewInt(62), big.NewInt(63), big.NewInt(64), big.NewInt(65), big.NewInt(66), big.NewInt(67), big.NewInt(68), big.NewInt(69)},
	{big.NewInt(70), big.NewInt(71), big.NewInt(72), big.NewInt(73), big.NewInt(74), big.NewInt(75), big.NewInt(76), big.NewInt(77), big.NewInt(78), big.NewInt(79)},
	{big.NewInt(80), big.NewInt(81), big.NewInt(82), big.NewInt(83), big.NewInt(84), big.NewInt(85), big.NewInt(86), big.NewInt(87), big.NewInt(88), big.NewInt(89)},
	{big.NewInt(90), big.NewInt(91), big.NewInt(92), big.NewInt(93), big.NewInt(94), big.NewInt(95), big.NewInt(96), big.NewInt(97), big.NewInt(98), big.NewInt(99)},
}

func main() {
	fmt.Printf("Matrix length: %d", len(matrix))
	routinesStart := time.Now()
	routinesProduct := routines(matrix)
	routinesExecTime := time.Since(routinesStart)

	fmt.Printf("Routines took %s\n", routinesExecTime)
	fmt.Printf("Routines result %v\n", routinesProduct)

	flatStart := time.Now()
	flatProduct := flat(matrix)
	flatExecTime := time.Since(flatStart)

	fmt.Printf("Flat took %s\n", flatExecTime)
	fmt.Printf("Flat result %v\n", flatProduct)
}

// routines is a function that multiplies a matrix using go routines
func routines(matrix [][]*big.Int) *big.Int {
	products := make(chan *big.Int, len(matrix))

	// start each goroutine, one array per routine
	for _, arr := range matrix {
		go func(arr []*big.Int) {
			product := multiplyArray(arr)
			products <- product
		}(arr)
	}

	// grab the values into one array to multiply them
	var productsArray []*big.Int
	for range matrix {
		productsArray = append(productsArray, <-products)
	}

	return multiplyArray(productsArray)
}

// flat is a function that multiplies a matrix without the use of go routines
func flat(matrix [][]*big.Int) *big.Int {
	var productsArray []*big.Int
	for _, arr := range matrix {
		productsArray = append(productsArray, multiplyArray(arr))
	}

	return multiplyArray(productsArray)
}
