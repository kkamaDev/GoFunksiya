// Kamron

// sharti  --- Berilgan n sonini raqamlarini teskari tartib yozishda xosil bo’lgan sonni topuvchi dastur tuzing

package main

import "fmt"

func main() {
	var n int = 10
	reverse(n)
}

func reverse(num int) {

	for i := num; i > 0; i-- {
		fmt.Println(i)
	}
}
