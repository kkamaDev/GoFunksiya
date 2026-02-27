// Kamron

// sharti = a, b va c sonlari berilgan. Bu sonlarning o’rtachasini topuvchi dastur tuzing!

package main

import "fmt"

func main() {
	fmt.Println(middle(16, 23, 66))

}

func middle(a, b, c int) int {
	var result int

	if a > b && a < c {
		result = a

	} else if b > a && b < c {
		result = b
	} else {
		result = c
	}
	return result
}
