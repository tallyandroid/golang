package main

import (
	"fmt"
)

var y string

func main() {
	//  DECLARE a variable to be of a certain type
	// and then ASSIGN A VALUE of that type to the variable

	fmt.Println("printing the value of y", y, "ending")
	fmt.Printf("%T", y)
	y = "Bond, James Bond"

	fmt.Println("printing the value of y", y, "ending")
	fmt.Printf("%T", y)

}