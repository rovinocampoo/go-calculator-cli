package main

import "fmt"

func add(a, b float64) float64 {
	sum := a + b
	return sum
}

func subtract(a, b float64) float64 {
	diff := a - b
	return diff
}

func multiply(a, b float64) float64 {
	prod := a * b
	return prod
}

func divide(a, b float64) (float64, bool) {
	if b != 0 {
		div := a / b
		return div, true
	}
	return 0, false
}

func main() {
	fmt.Println("Hello, Welcome to Rovin's first Go project. The Mighty Calculator!")

	var firstNum float64
	var secondNum float64

	fmt.Print("Enter First Number:")
	fmt.Scan(&firstNum)
	fmt.Print("Enter Second Number:")
	fmt.Scan(&secondNum)

	fmt.Printf("Sum: %.2f\n", add(firstNum, secondNum))
	fmt.Printf("Difference: %.2f\n", subtract(firstNum, secondNum))
	fmt.Printf("Product: %.2f\n", multiply(firstNum, secondNum))

	result, ok := divide(firstNum, secondNum)
	if ok {
		fmt.Printf("Quotient: %.2f\n", result)
	} else {
		fmt.Println("Quotient: Can't divide by zero good sir.")
	}

}
