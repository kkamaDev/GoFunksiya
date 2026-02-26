// Kamron

// sharti = To’rt xonali a soni berilgan! Jumlani rostilkka tekshiring! A sonining juft o’rindagi raqamlari yigindisi va
// toq o’rindiadi ramlari yig’indisining ayrimasi 0 ga teng va a soning xar-bir raqami takrorlanmas

package main

import "fmt"

func main() {
	Funksiya(7833)

}
func Funksiya(a int) {
	t1 := a / 1000
	j1 := a / 100 % 10
	t2 := a / 10 % 10
	j2 := a % 10
	natija := (t1+t2) == (j1+j2) && t1 != j1 && t1 != t2 && t1 != j2 && j1 != t2 && t2 != j2
	fmt.Println(natija)

}
