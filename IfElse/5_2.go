// Kamron
// sharti = a soni berilgan! Agar a soni 3 ga va 5 ga bo’linsa FIZZBUZZ, agar faqat 3 ga bo’linsa FIZZ,
//  agar faqat 5 ga bo’linssa BUZZ so’zini ekranga chiqaruchi dastur tuzing!

package main

import "fmt"

func main() {
	fmt.Println(fizzbuzz(65))

}

func fizzbuzz(son int) string {
	var natija string
	if son%3 == 0 && son%5 == 0 {
		natija = "FIZZBUZZ"
	} else if son%3 == 0 {
		natija = "FIZZ"
	} else if son%5 == 0 {
		natija = "BUZZ"
	}
	return natija
}
