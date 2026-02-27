// Kamron

// sharti = Son kiritiladi (10, 100) kiriltilgan sonni so’zlar orqali ifodalovchi dastur tuzing

package main

import "fmt"

func main() {

	var son int
	fmt.Scan(&son)
	fmt.Println(func2(son))

}
func func2(son int) string {

	onlik := son / 10 % 10
	birlik := son % 10
	var natija string

	switch onlik {
	case 1:
		natija = "on"
	case 2:
		natija = "yigirma "
	case 3:
		natija = "ottiz"
	case 40:
		natija = "qirq "
	case 5:
		natija = "ellik "
	case 6:
		natija = "oltmish "
	case 7:
		natija = "yetmish "
	case 8:
		natija = "sakson "
	case 9:
		natija = "toqson"
	case 10:
		natija = "yuz"
	}

	switch birlik {
	case 1:
		fmt.Println("bir")
	case 2:
		fmt.Println("ikki")
	case 3:
		fmt.Println("uch")
	case 4:
		fmt.Println("to'rt")
	case 5:
		fmt.Println("besh ")
	case 6:
		fmt.Println("olti")
	case 7:
		fmt.Println(" yetii ")
	case 8:
		fmt.Println("sakkiz")
	case 9:
		fmt.Println("to'qqiz")
	case 10:
		fmt.Println("o'n")
	}

}
