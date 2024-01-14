package main

import (
	"fmt"
	"mypack/channel"
)

func main() {
	// condition.Biggest1()
	// condition.Biggest2()
	// condition.Biggest3()
	// condition.Biggest4()

	//loops.Guess()
	//loops.Guess2()
	//loops.Guess3()
	//loops.Guess4()
	//loops.Guess5()

	// arrays.Array1()
	// arrays.Array2()
	// arrays.Array3()
	// arrays.Array4()

	//slices.Slice1()

	//sum := functions.Sum(10, 20)
	// fmt.Println(sum)

	//sum , sub, _, div := functions.Calculate(10, 20)
	//fmt.Println(sum, sub, div)

	//fmt.Println(functions.SumVariadic(10, 20, 30, 40, 50))

	//numbers := []int{10, 20, 30, 40, 50}
	//fmt.Println(functions.SumVariadic2(numbers...))

	//maps.Map1()

	//ranges.Range1()
	//ranges.OddEven()

	//ranges.Range2()

	//num := 10
	//pointers.Pointers1(&num)

	//structs.Struct1()
	
	// channel 
	evenCh := make(chan int)
	oddCh := make(chan int)

	go channel.Channel1(evenCh)
	go channel.Channel2(oddCh)

	even, odd := <-evenCh, <-oddCh

	mul := even * odd

	fmt.Println("Even: ", even)
	fmt.Println("Odd: ", odd)
	fmt.Println("Multiplication: ", mul)

}
