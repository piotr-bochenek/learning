package main

import "fmt"

const pi = 3.1415
const (
	first  = 1
	second = "second"
	third  = iota
	_
	_
	fourth
)

func main() {
	const constValue int = 123
	fmt.Println(constValue + 3)

	fmt.Println(float32(constValue) + 1.2)

	fmt.Println(pi, first, second, third, fourth)
}
