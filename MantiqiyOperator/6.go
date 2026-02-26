// Kamron
// sharti a va b sonlari berilgan! Bu sonlarning yig’indisi musbatligini va toqligini aniqlovchi dastur

package main

import "fmt"

func main() {
	toqANDMusbat(14, 3)

}

func toqANDMusbat(a, b int) {
	fmt.Println(a+b > 0 && a+b%2 != 0)
}
