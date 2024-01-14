package condition

import (
	"fmt"
)

func Biggest1() {

	a, b, c := 10, 20, 30
	if a > b && a > c {
		fmt.Println("a en büyük sayidir: ", a)
	}

	if b > a && b > c {
		fmt.Println("b en büyük sayidir: ", b)
	}

	if c > a && c > b {
		fmt.Println("c en büyük sayidir.", c)
	}

}

func Biggest2() {
	a, b, c := 10, 20, 30
	if a > b && a > c {
		fmt.Println("a en büyük sayidir: ", a)
	} else if b > a && b > c {
		fmt.Println("b en büyük sayidir: ", b)
	} else if c > a && c > b {
		fmt.Println("c en büyük sayidir.", c)
	}
}

func Biggest3() {
	a, b, c := 10, 20, 30
	if a > b && a > c {
		fmt.Println("a en büyük sayidir: ", a)
	} else if b > a && b > c {
		fmt.Println("b en büyük sayidir: ", b)
	} else {
		fmt.Println("c en büyük sayidir.", c)
	}
}

func Biggest4() {
	a, b, c := 10, 20, 30
	biggest := a
	if b > biggest {
		biggest = b
	}
	if c > biggest {
		biggest = c
	}
	fmt.Println("biggest: ", biggest)

}
