// Kamron

// Berilgan a soni musbatligini aniqlovchi dastur tuzing

package main

import "fmt"

func main() {
	result := musbat(-45)
	fmt.Println("Musbatmi : ", result)

}

func musbat(num int) bool {
	Ismusbat := num > 0
	return Ismusbat
}
