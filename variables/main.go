package main

import (
	"fmt"
)

func main() {
	var myIntVar int
	myIntVar = -12
	fmt.Println("My variable is: ", myIntVar)

	var myUintVar uint
	myUintVar = 12
	fmt.Println("My variable is: ", myUintVar)

	var myStringVar string
	myStringVar = "My String variable"
	fmt.Println("My variable is: ", myStringVar)

	var myBoolVar bool
	myBoolVar = true
	fmt.Println("My variable is: ", myBoolVar)

	fmt.Println("My variable address is: ", &myStringVar)

	myIntVar2 := 23

	fmt.Println("My int variable with :=", myIntVar2)

	myString2 := "My string variable with :="
	fmt.Println(myString2)

	const myFirstConst = "a12"
	fmt.Println("My const is: ", myFirstConst)

	const myIntConst int = 12
	fmt.Println("My const is: ", myIntConst)

	const myStringConst string = "aaa"
	fmt.Println("My const is: ", myStringConst)
}
