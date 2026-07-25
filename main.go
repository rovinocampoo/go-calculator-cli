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

	sum := firstNum+secondNum
	diff := firstNum-secondNum
	mult := firstNum * secondNum
	div := firstNum/secondNum

	fmt.Println("Sum:" ,sum)
	fmt.Println("Difference:" ,diff)
	fmt.Println("Product:" ,mult)
	fmt.Println("Quotient:" ,div)



}
