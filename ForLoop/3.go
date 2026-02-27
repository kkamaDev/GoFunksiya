// Kamron

// sharti - 1 dan 10 gacha bo’lgan sonlarni yig’indisni ekranga chiqaruchi dastur tuzing

package main

import "fmt"

func main() {
	sum(10)

}

func sum(x int) {
	sum := 0

	for i := 0; i < x; i++ {
		sum += i

	}
	fmt.Println(sum)
}
