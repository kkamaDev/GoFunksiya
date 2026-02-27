// Kamron
// sharti = a va b sonlari berilgan. Bu sonlarning kattasini topuvchi dastur tuzing!

package main

import "fmt"

func main() {

	AKATTAB(1, 8)

}

func AKATTAB(a, b int) {

	if a > b {
		fmt.Println(a, ">", b)
	} else {
		fmt.Println(b, ">", a)
	}
}
