// Kamron
// sharti = 1 dan 10 gacha bo’lgan sonlarni teskari tartibda ekranga chiqaruchi dastur tuzing

package main

import "fmt"

func main() {
	Loop(10)

}

func Loop(x int) {
	for i := x; i > 0; i-- {

		fmt.Println(i)

	}
}
