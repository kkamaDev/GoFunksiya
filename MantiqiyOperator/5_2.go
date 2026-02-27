// Kamron
// sharti a va b sonlarni berilgan jumlani rostlikka tekshiring: a soni b sonidan katta yoki teng

package main

import "fmt"

func main() {
	natija := Katta(8, 7)
	fmt.Println(natija)
}

func Katta(a, b int) bool {
	natija := a >= b
	return natija
}
