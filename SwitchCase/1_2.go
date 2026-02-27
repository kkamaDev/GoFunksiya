// Kamron
// sharti: Son kiritiladi (1, 10) kiriltilgan sonni so’zlar orqali ifodalovchi dastur tuzing

package main

import "fmt"

func main() {
	var num int
	fmt.Scan(&num)
	fmt.Println(switch1(num))

}

func switch1(son int) string {
	var natija string
	switch son {
	case 1:
		natija = "bir"
	case 2:

		natija = "ikki"
	case 3:
		natija = "uch "
	case 4:
		natija = "to'rt"
	case 5:
		natija = "besh"
	case 6:
		natija = "olti"
	case 7:
		natija = "yetti"
	case 8:
		natija = "sakkiz"
	case 9:
		natija = "to'qqiz"
	case 10:
		natija = "o'n"

	}
	return natija

}
