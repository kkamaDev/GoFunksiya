// Kamron

// sharti - 1 dan 10 gacha bo’lgan sonlarni yig’indisni ekranga chiqaruchi dastur tuzing

package main

import "fmt"

func main() {
	fmt.Println(Sum(10))

}

func Sum(son int) int {
	summa := 0
	for i := 0; i < son; i++ {
		summa += i
	}
	return summa
}
