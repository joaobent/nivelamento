package main
import (
	"fmt"
)

func main(){
	a:= 100==100
	b:= 10!=100
	c:= 10<100
	d:= 10>100
	e:= 10<=100
	f:= 10>=100
	fmt.Printf("%v\n%v\n%v\n%v\n%v\n%v", a, b, c, d, e, f)
	
}