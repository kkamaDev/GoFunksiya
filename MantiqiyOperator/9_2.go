// Kamron
// A soni berilgan! Jumlani rostlikka tekshiring! A soni 3 xonali va toq son

package main

import "fmt"

func main() {
	fmt.Println(three(55))

}
func three(A int) bool {
	Onlar := A / 100

	natija := Onlar > 0 && A%2 != 0
	return natija

}
