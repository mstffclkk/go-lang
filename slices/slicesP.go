package slices

import (
	"fmt"
)

func Slice1() {

	num := []int{1, 2, 3, 4, 5}
	fmt.Println(num)

	num2 := make([]int, 5)
	copy(num2, num)

	fmt.Println(num2)

}
