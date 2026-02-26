// Kamron

// sharti = a va b sonlarni berilgan jumlani rostlikka tekshiring: a soni b sonidan katta

package main

import "fmt"

func main() {
	natija := bigger(8, 7)
	fmt.Println(natija)
}

func bigger(a, b int) bool {
	natija := a > b
	return natija
}
