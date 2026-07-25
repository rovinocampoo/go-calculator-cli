package main

import "fmt"

func main() {
	fmt.Println("Hello, Welcome to Rovin's first Go project. The Mighty Calculator!")

	var firstNum float64
	var secondNum float64

	fmt.Print("Enter First Number:")
	fmt.Scan(&firstNum)
	fmt.Print("Enter Second Number:")
	fmt.Scan(&secondNum)

	sum := firstNum + secondNum
	diff := firstNum - secondNum
	mult := firstNum * secondNum
	div := firstNum / secondNum

	fmt.Printf("Sum: %.2f\n", sum)
	fmt.Printf("Difference: %.2f\n", diff)
	fmt.Printf("Product: %.2f\n", mult)
	if secondNum != 0 {
		fmt.Printf("Quotient: %.2f\n", div)
	} else {
		fmt.Println("Can't divide by zero good sir.")
	}

}
