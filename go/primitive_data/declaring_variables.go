package main

import "fmt"

func main() {
	var i int
	i = 1
	fmt.Println(i)

	var f float32 = 1.23
	fmt.Println(f)

	stringValue := "String"
	fmt.Println(stringValue)

	b := true
	fmt.Println(b)

	c := complex(3, 4)
	fmt.Println(c)

	r, im := real(c), imag(c)
	fmt.Println(r, im)
}
