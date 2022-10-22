package main

func add(a, b int) int {
	return a + b
}
func fact(num int) int {
	result := 1
	for i := 1; i <= num; i++ {
		result *= i
	}
	return result
}
