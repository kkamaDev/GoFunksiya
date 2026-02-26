//  Kamron

//  sharti = // kki xonal a soni berilgan! Jumlani rostlikka tekshiring: a soni polindrom va toq son

package main

import "fmt"

func main() {
	fmt.Println(ispolindrome(88))

}

func ispolindrome(a int) bool {

	result := a%11 == 0 && a%2 != 0
	return result
}
