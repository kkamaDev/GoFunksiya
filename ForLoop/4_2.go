// Kamron

// sharti = 1 dan 10 gacha bo’lgan sonlar ichidan toqlarini ekranga chiqaruchi dastur tuzing!

package main

import (
	"fmt"
	"strconv"
)

func main() {

	fmt.Println(toq(10))

}

func toq(num int) string {
	var natija string

	for i := 0; i < num; i++ {
		if i%2 != 0 {
			natija += strconv.Itoa(i) + " "
		}

	}
	return natija

}
