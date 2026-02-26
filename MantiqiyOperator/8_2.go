// Kamron

// sharti Uch xonali a soni berilgan jumlani rostlikka tekshiring: a soni polindrom va juft son

package main

import "fmt"

func main() {
	fmt.Println(juftPolindrom(303))

}

func juftPolindrom(a int) bool {

	yuzlar := a / 100
	birlar := a % 10
	natija := yuzlar == birlar && a%2 == 0
	return natija
}
