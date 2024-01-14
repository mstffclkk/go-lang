package maps

import "fmt"

func Map1() {
	ages := map[string]int{
		"John":  32,
		"Peter": 28,
		"Mary":  25,
	}

	fmt.Println(ages)

	value, ok := ages["John"]
	fmt.Println("Value: ", value, "Present?", ok)

}