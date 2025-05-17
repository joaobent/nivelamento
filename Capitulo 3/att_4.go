package main

import (
	"fmt"
)

type bola int
var x bola

func main(){
	fmt.Println(x)
	fmt.Printf("%T\n", x)
	x=42
	fmt.Println(x)

}