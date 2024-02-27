package main

func Sum(array []int) int {
	result := 0
	for _, v := range array {
		result += v
	}
	return result
}