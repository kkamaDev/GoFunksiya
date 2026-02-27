// Kamron
// sharti = a, b va c sonlari berilgan. Bu sonlarning eng kichigini topuvchi dastur tuzing!

package main

import "fmt"

func main() {
	fmt.Println("eng kichik son", Smallest(6, 8, 7))

}

func Smallest(a, b, c int) int {
	var natija int
	if a < b && a < c {
		natija = a
	} else if b < a && b < c {
		natija = b
	} else {
		natija = c
	}
	return natija
}
