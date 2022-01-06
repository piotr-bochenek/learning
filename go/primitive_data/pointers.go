package main

import "fmt"

func main() {
	var firstValue *string = new(string)
	*firstValue = "value"
	fmt.Println(*firstValue)

	secondValue := "value"

	secondValuePointer := &secondValue

	fmt.Println(secondValue, *secondValuePointer)
	fmt.Println(secondValuePointer, *secondValuePointer)

	secondValue = "changed_value"
	fmt.Println(secondValuePointer, *secondValuePointer)
}
