package main

import "fmt"

func main() {
	x := 15
	a := &x // memory address with & being the "pointer"
	y := 20
	b := &y // memory address pointing to y via b
	fmt.Println(a)
	//fmt.Println(*a)
	//*a = 5
	//fmt.Println(x)
	fmt.Println(b)
	fmt.Println(x+y)

	//the * aka the asterisk/star allows you to READ THROUGH
	//a memory address. so...

	fmt.Println(*b + *a)
}
