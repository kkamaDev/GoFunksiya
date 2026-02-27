// Kamronbek

// sharti = 1 dan 10 gacha bo’lgan sonlarni kvadratlarini ekranga chiqaruchi dastur tuzing

package main

import "fmt"

func main() {
	Kvardat(10)

}

func Kvardat(num int) {

	for i := 0; i < num; i++ {
		fmt.Println(i * i)
	}
}
