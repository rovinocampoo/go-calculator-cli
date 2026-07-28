package main

import (
	"bufio"
	"fmt"
	"os"
)

func readNumbers() (float64, float64) {
	var firstNum float64
	var secondNum float64
	fmt.Print("Enter First Number:")
	fmt.Scan(&firstNum)
	fmt.Print("Enter Second Number:")
	fmt.Scan(&secondNum)
	return firstNum, secondNum
}
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
			reader := bufio.NewReader(os.Stdin)

	for {

		var choice int
		fmt.Println("==================================================================")
		fmt.Println("Hello, Welcome to Rovin's first Go project. The Mighty Calculator!")
		fmt.Println("==================================================================")

		fmt.Println("\n1. Add")
		fmt.Println("2. Subtract")
		fmt.Println("3. Multiply")
		fmt.Println("4. Divide")
		fmt.Println("5. Exit")

		fmt.Print("Choose (1-5):")
		fmt.Scan(&choice)
		if choice == 5 {
			fmt.Println("Thanks for using Rovin's Mighty Calculator!")
			break
		}
		if choice < 1 || choice > 4 {
			fmt.Println("Invalid number.")
			continue
		}

		firstNum, secondNum := readNumbers()
		switch choice {
		case 1:
			fmt.Printf("Sum: %.2f\n", add(firstNum, secondNum))
		case 2:
			fmt.Printf("Difference: %.2f\n", subtract(firstNum, secondNum))
		case 3:
			fmt.Printf("Product: %.2f\n", multiply(firstNum, secondNum))
		case 4:
			result, ok := divide(firstNum, secondNum)
			if ok {
				fmt.Printf("Quotient: %.2f\n", result)
			} else {
				fmt.Println("Quotient: Can't divide by zero good sir.")
			}
		}
		fmt.Print("Press Enter to Continue...")
		_, _ = reader.ReadBytes('\n')
		_, _ = reader.ReadBytes('\n')

	}

}
