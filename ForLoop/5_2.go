// 1 dan 10 gacha bo’lgan sonlar ichidan juftlarini ekranga chiqaruchi dastur tuzing!
// kamron

package main

import (
	"fmt"
	"strconv"
)

func main() {

	fmt.Println(juft(10))

}

func juft(num int) string {
	var natija string

	for i := 0; i < num; i++ {
		if i%2 == 0 {
			natija += strconv.Itoa(i) + " "
		}

	}
	return natija

}
