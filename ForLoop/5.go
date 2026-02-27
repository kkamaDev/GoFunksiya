// 1 dan 10 gacha bo’lgan sonlar ichidan juftlarini ekranga chiqaruchi dastur tuzing!
// kamron

package main

import "fmt"

func main() {
	Juft(10)

}

func Juft(num int) {

	for i := 0; i < num; i++ {
		if i%2 == 0 {
			fmt.Println(i)
		}
	}
}
