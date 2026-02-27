// Kamron

// sharti = a soni berilgan! Agar a soni 3 ga va 5 ga bo’linsa FIZZBUZZ, agar faqat 3 ga bo’linsa FIZZ,
//  agar faqat 5 ga bo’linssa BUZZ so’zini ekranga chiqaruchi dastur tuzing!

package main

import "fmt"

func main() {
	FizzBUzz(15)

}

func FizzBUzz(son int) {
	if son%3 == 0 && son%5 == 0 {
		fmt.Println("FIZZBUZZ")
	} else if son%3 == 0 {
		fmt.Println("FIZZ")
	} else if son%5 == 0 {
		fmt.Println("BUZZ")
	}
}
