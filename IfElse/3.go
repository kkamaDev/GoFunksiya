// Kamron

// sharti = a, b va c sonlari berilgan. Bu sonlarning o’rtachasini topuvchi dastur tuzing!

package main

import "fmt"

func main() {
	Middle(1, 4, 3)

}

func Middle(a, b, c int) {
	if a > b && a < c {
		fmt.Println(a, "o'rtacha son ")
	} else if b > a && b < c {
		fmt.Println(b, "o'rtacha son ")
	} else {
		fmt.Println(c, "o'rtacha son ")
	}
}
