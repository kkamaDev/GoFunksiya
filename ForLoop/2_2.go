// Kamronbek

// sharti = 1 dan 10 gacha bo’lgan sonlarni kvadratlarini ekranga chiqaruchi dastur tuzing

package main

import (
	"fmt"
	"strconv"
)

func main() {
	fmt.Println(kvadrat(10))

}

func kvadrat(num int) string {
	var natija string

	for i := 0; i < num; i++ {

		result := i * i
		natija += strconv.Itoa(result) + " "
	}
	return natija
}
