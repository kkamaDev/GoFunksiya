// Kamron

// sharti - Berilgan a soni juftligini aniqlovchi dastur tuzing

package main

import "fmt"

func main() {
	natija := ISJuft(45)
	fmt.Println("son juftmi : ", natija)

}

func ISJuft(x int) bool {

	result := x%2 == 0
	return result

}
