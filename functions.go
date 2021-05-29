package main

import "fmt"

func add1(x int, y int) (int, int) {
	var result = x + y
	result1 := x - y
	return result, result1
}

func main() {
	num1 := 5
	num2 := 7

	out1, out2 := add1(num1, num2)
	fmt.Println(out1, out2)
}
