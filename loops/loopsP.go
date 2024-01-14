package loops

import (
	"fmt"

)

func Guess() {
	num := 10
	var guess int

	for {
		fmt.Println("Enter your guess: ")
		fmt.Scanln(&guess)
		// fmt.Println("Your guess is: ", guess)
		if guess == num {
			fmt.Println("Your guess is True!")
			break
		} else if guess < num {
			fmt.Println("Your guess is less than num. Try again!")
		} else {
			fmt.Println("Your guess is bigger than num. Try again!")
		}
	}
}

func Guess2() {
	num := 10
	var guess int

	counter := 0

	for {
		fmt.Println("Enter your guess: ")
		fmt.Scanln(&guess)
		// fmt.Println("Your guess is: ", guess)
		if guess == num {
			fmt.Println("Your guess is True!")
			counter++
			break
		} else if guess < num {
			fmt.Println("Your guess is less than num. Try again!")
			counter++
		} else {
			fmt.Println("Your guess is bigger than num. Try again!")
			counter++

		}
	}

	fmt.Println("You tried ", counter, " times.")

	if (0 < counter) && (counter < 3) {
		fmt.Println("You are a Wizard!")
	} else if (3 <= counter) && (counter < 6) {
		fmt.Println("You are a good guesser!")
	} else if (6 <= counter) && (counter < 10) {
		fmt.Println("You are a average guesser!")
	} else {
		fmt.Println("You are a bad guesser!")
	}

}

func Guess3() {

	num := 10
	var guess int

	for {
		fmt.Println("Enter your guess: ")
		fmt.Scanln(&guess)
		// fmt.Println("Your guess is: ", guess)
		switch {
		case guess == num:
			fmt.Println("Your guess is True!")
			return
		case guess < num:
			fmt.Println("Your guess is less than num. Try again!")
		case guess > num:
			fmt.Println("Your guess is bigger than num. Try again!")
		}
	}
}

func Guess4() {
	// prime numbers

	var num int
	var i int

	fmt.Println("Enter a number: ")
	fmt.Scanln(&num)

	for i = 2; i < num; i++ {
		if num%i == 0 {
			fmt.Println("Not a prime number!")
			break
		}
	}

	if i == num {
		fmt.Println("Prime number!")
	}

}

func Guess5() {
	// amicable numbers

	num1 := 0
	num2 := 0

	sum1 := 0
	sum2 := 0

	fmt.Println("Enter 2 numbers: ")
	fmt.Scanln(&num1)
	fmt.Scanln(&num2)



	for i := 1; i < num1; i++ {
		if num1%i == 0 {
			//fmt.Println(i)
			sum1 += i
		}

	}

	fmt.Println("num1 sum: ", sum1)

	for i := 1; i < num2; i++ {
		if num2%i == 0 {
			//fmt.Println(i)
			sum2 += i
		}
		
	}

	fmt.Println("num2 sum: ", sum2)

	if (sum1 == num2) && (sum2 == num1) {
		fmt.Println("Amicable numbers!")
	} else {
		fmt.Println("Not amicable numbers!")
	}
}	

