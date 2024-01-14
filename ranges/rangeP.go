package ranges

import "fmt"

func Range1() {
	numbers := []int{10, 20, 30, 40, 50}

	for i, number := range numbers {
		fmt.Println("Index: ", i, "Number: ", number)
		fmt.Println(number)
	}

}

func OddEven() {
	numbers := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	oddSum := 0
	evenSum := 0

	for _, number := range numbers {
		if number%2 == 0 {
			fmt.Println(number, "is even")
			evenSum += number
		} else {
			fmt.Println(number, "is odd")
			oddSum += number
		}
	}

	fmt.Println("Odd sum: ", oddSum)
	fmt.Println("Even sum: ", evenSum)
}


func Range2() {
	ages := map[string]int{
		"John":  32,
		"Peter": 28,
		"Mary":  25,
	}

	for k, v := range ages {
		fmt.Println(k, "is", v, "years old")
	}
}