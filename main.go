package main

import "fmt"

func main() {

	var operator = "+"
	var num1 float32
	var num2 float32

	num1 = 1
	num2 = 5

	result, err := calculator(num1, num2, operator)
	if err != nil {
		fmt.Println(err.Error())
		return
	}
	fmt.Printf("result = %f\n", result)
}

func calculator(num1 float32, num2 float32, operator string) (result float32, err error) {
	if operator == "+" {
		result = num1 + num2
	} else if operator == "-" {
		result = num1 - num2
	} else if operator == "*" {
		result = num1 * num2
	} else if operator == "/" {
		if num1 == 0 {
			return 0, fmt.Errorf("error divide by zero")
		}
		result = num1 / num2
	} else {
		return 0, fmt.Errorf("this operator not support")
	}

	return
}
