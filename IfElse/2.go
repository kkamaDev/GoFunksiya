// Kamron
// sharti = a, b va c sonlari berilgan. Bu sonlarning eng kattasini topuvchi dastur tuzing!

package main

import "fmt"

func main() {
	Compire(1, 3, 5)

}

func Compire(a, b, c int) {

	if a > b && a > c {
		fmt.Println(a, "eng katta son ")
	} else if b > a && b > c {
		fmt.Println(b, "eng katta son ")
	} else {
		fmt.Println(c, "eng katta son")
	}
}
