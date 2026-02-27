// Kamron
// sharti = a, b va c sonlari berilgan. Bu sonlarning eng kattasini topuvchi dastur tuzing!

package main

import "fmt"

func main() {
	fmt.Println(compire(1, 2, 3))

}

func compire(a, b, c int) int {
	var result int
	if a > b && a > c {
		result = a

	} else if b > a && b > c {
		result = b
	} else {
		result = c
	}
	return result
}
