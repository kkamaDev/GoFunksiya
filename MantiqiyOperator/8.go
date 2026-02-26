// Kamron

// sharti Uch xonali a soni berilgan jumlani rostlikka tekshiring: a soni polindrom va juft son

package main

import "fmt"

func main() {
	JufTpolindrom(101)

}

func JufTpolindrom(a int) {

	yuzlar := a / 100
	birlar := a % 10
	fmt.Println(yuzlar == birlar && a%2 == 0)
}
