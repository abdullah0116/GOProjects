package main

import (
	"fmt"
	"math"
)

func main() {
	var firstNumber float64
	var opperation string  
	var secondNumber float64

	fmt.Print("please enter the first number you'd like to add: ")
	fmt.Scan(&firstNumber)
	
	fmt.Print("please enter the first number you'd like to add: ")
	fmt.Scan(&opperation)

	fmt.Print("please enter the first number you'd like to add: ")
	fmt.Scan(&secondNumber)

	if (opperation == "+") {
		fmt.Println("Here's your answer: ", firstNumber + secondNumber)
	} else if (opperation == "-") {
		fmt.Println("Here's your answer: ", firstNumber - secondNumber)
	} else if (opperation == "*") {
		fmt.Println("Here's your answer: ", firstNumber * secondNumber)
	} else if (opperation == "/") {
		fmt.Println("Here's your answer: ", firstNumber / secondNumber)
	} else if (opperation == "%") {
		fmt.Println("Here's your answer: ", math.Mod(firstNumber, secondNumber))
	} else {
		fmt.Println("Sorry cant find this sign: ", opperation)
	}
}
