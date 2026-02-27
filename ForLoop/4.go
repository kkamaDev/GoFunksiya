// Kamron

// sharti = 1 dan 10 gacha bo’lgan sonlar ichidan toqlarini ekranga chiqaruchi dastur tuzing!

package main

import "fmt"

func main() {
	Toq(10)

}

func Toq(num int) {

	for i := 0; i < num; i++ {
		if i%2 != 0 {
			fmt.Println(i)
		}
	}
}
