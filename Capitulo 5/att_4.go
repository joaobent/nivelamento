package main

import "fmt"



func main() {

	x := 10
	fmt.Printf("%d %b %#x\n", x, x, x)
	a := x<<1
	fmt.Printf("%d %b %#x\n", a, a, a)


}