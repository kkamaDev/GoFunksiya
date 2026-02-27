// Kamron
// sharti = a, b va c sonlari berilgan. Bu sonlarning eng kichigini topuvchi dastur tuzing!

package main

import "fmt"

func main() {
	smallest(9, 8, 4)

}
func smallest(a, b, c int) {

	if a < b && a < c {
		fmt.Println(a, "eng kichik son ")
	} else if b < a && b < c {
		fmt.Println(b, "eng kichik son ")
	} else {
		fmt.Println(c, "eng kichik son ")
	}
}
