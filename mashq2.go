package main

import "fmt"

func main() {

	natija := func(x, y int) int {
		return x + y
	}
	natija2 := func(x, y int) int {
		return x + y
	}

	f := FuncPractis(natija, natija2)
	fmt.Println(f(0, 0))

}

func FuncPractis(fn1 func(x, y int) int, fn2 func(x, y int) int) func(x, y int) int {
	result := fn1(10, 4)
	result2 := fn2(1, 3)

	return func(x, y int) int {
		if (result+result2)%2 == 0 {
			return result - result2
		}
		return result + result2
	}

}
