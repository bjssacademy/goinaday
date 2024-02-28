package main

func Sum(x int, y int) int {
	return x + y
}

func ComplexSum(array []int) int {
	result := 0
	for _, v := range array {
		result += v
	}
	return result
}