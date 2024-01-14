package arrays

import (
	"fmt"
)

func Array1() {
	// biggest number in array
	arr1 := [5]int{1, 2, 3, 4, 5}
	biggest := arr1[0]

	for i := 0; i < len(arr1); i++ {
		if arr1[i] > biggest {
			biggest = arr1[i]
		}

	}

	fmt.Println("Biggest: ", biggest)
}

func Array2() {
	// biggest number in array
	arr1 := [5]int{1, 2, 3, 4, 5}
	biggest := arr1[0]

	for _, v := range arr1 {
		if v > biggest {
			biggest = v
		}

	}

	fmt.Println("Biggest: ", biggest)
}

func Array3() {
	// 2D array
	arr1 := [2][3]int{{1, 2, 3}, {4, 5, 6}}
	
	for i := 0; i < len(arr1); i++ {
		for j := 0; j < len(arr1[i]); j++ {
			fmt.Print(arr1[i][j], " ")
		}
		fmt.Println()
	}
}

func Array4() {

	arr:= []int{1,2,3,4,5}

	arr = append(arr, 6)

	fmt.Println(arr)

}