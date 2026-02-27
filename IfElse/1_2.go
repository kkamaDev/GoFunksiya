// Kamron
// sharti = a va b sonlari berilgan. Bu sonlarning kattasini topuvchi dastur tuzing!

package main

import "fmt"

func main() {
	fmt.Println(akattab(4, 6))

}
func akattab(a, b int) bool {

	var taqqoslash bool
	if a > b {
		taqqoslash = true
	} else {
		taqqoslash = false
	}
	return taqqoslash

}
