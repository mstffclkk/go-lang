package functions

func Sum(number1, number2 int) int {
	return number1 + number2

}

func Calculate(number1, number2 int) (int, int, int, float64) {
	sum := number1 + number2
	sub := number1 - number2
	mul := number1 * number2
	div := float64(number1) / float64(number2)

	return sum, sub, mul, div
}

func SumVariadic(numbers ...int) int {
	sum := 0
	for _, value := range numbers {
		sum += value
	}
	return sum
}

func SumVariadic2(numbers ...int) int {
	sum := 0
	for i := 0; i < len(numbers); i++ {
		sum += numbers[i]
	}
	return sum
}
