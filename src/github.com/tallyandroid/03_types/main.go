package main

import 
(
	"fmt"
)

var y = 42 //Holds value of type int

//  DECLARE the VARIABLE with the IDENTIFIER "z"
//  is of TYPE string
//  and I ASSIGN the VALUE "shaken, not stirred"

var z = "Shaken, not stirred"
var a string = `James said, "Shaken, not stirred"`

func main() {
	fmt.Println(y)
	fmt.Printf("%T\n", y)
	fmt.Println(z)
	fmt.Printf("%T\n", z)
	fmt.Println(a)
	fmt.Printf("%T\n", a)
}