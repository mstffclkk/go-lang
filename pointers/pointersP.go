package pointers

import "fmt"

func Pointers1(num *int) {
	*num = *num + 1

	fmt.Println("Value inside function: ", *num)
	fmt.Println("Address inside function: ", num)

}
