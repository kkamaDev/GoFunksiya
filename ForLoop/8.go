// Kamronbek

// sharti - Berilgan n sonini tub yoki tub emasligini aniqlovchi dastur tuzing!

package main

import "fmt"

func main() {
	Tub(10)

}

func Tub(num int) {

	for i := 1; i < num; i++ {
		if i%i == 0 && i%1 == 0 {
			fmt.Println(i, "tub son ")
		} else {
			fmt.Println(i, "tub emas ")
		}
	}
}
