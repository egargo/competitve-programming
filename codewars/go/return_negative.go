// https://www.codewars.com/kata/55685cd7ad70877c23000102/train/go

package main

import "fmt"

func MakeNegative(x int) int {
	if x > 0 {
		return x * -1
	}

	return x
}

func main() {
	fmt.Println(MakeNegative(12))
	fmt.Println(MakeNegative(-12))
	fmt.Println(MakeNegative(0))
}
