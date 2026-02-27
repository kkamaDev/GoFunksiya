// kamron

// sharti -- Berilgan n sonini poindrom yoki polindrom emasligini aniqlovchi dastur tuzing!

package main

import "fmt"

func main() {

	polindromm(861)

}

func polindromm(num int) {

	yuzlar := num / 100
	birlar := num % 10

	for i := 0; i < num; i++ {
		if i%11 == 0 {
			fmt.Println(i, "polindrome")
		} else if yuzlar == birlar {
			fmt.Println(i, "polindrome")
		}
	}
}
