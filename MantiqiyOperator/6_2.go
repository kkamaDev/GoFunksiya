// Kamron
// sharti a va b sonlari berilgan! Bu sonlarning yig’indisi musbatligini va toqligini aniqlovchi dastur

package main

import "fmt"

func main() {
	fmt.Println(ToqMUsbat(11, 1))

}

func ToqMUsbat(a, b int) bool {
	result := a+b > 0 && a+b%2 != 0
	return result
}
