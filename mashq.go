package main

import (
	"errors"
	"fmt"
)

func main() {

	result1, result2, error := DiveReturn(10, 3)
	if error != nil {
		return
	}
	fmt.Println(result1, result2, error)

}

func DiveReturn(a, b int) (int, int, error) {

	if b == 0 {
		return 0, 0, errors.New("0 ga bo'lib bo'lmaydi ")
	}
	result, result2 := a/b, a%b

	return result, result2, nil
}
