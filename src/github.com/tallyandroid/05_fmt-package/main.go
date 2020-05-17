package main


import (
	"fmt"
)

var y = 42

func main() {
	fmt.Println(y)
	fmt.Printf("%T\n", y)
	fmt.Printf("%b\n", y) //binary
	fmt.Printf("%x\n", y) //hexidecimal
	fmt.Printf("%#x\n", y)//hexidecimal with the 0 in front



}