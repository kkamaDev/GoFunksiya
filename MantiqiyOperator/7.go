//  Kamron

//  sharti = // kki xonal a soni berilgan! Jumlani rostlikka tekshiring: a soni polindrom va toq son

package main

import "fmt"

func main() {
	IsPolindrom(77)

}

func IsPolindrom(a int) {

	fmt.Println(a%11 == 0 && a%2 != 0)
}
