package main

import "fmt"

func main() {
	license := true
	age := 19

	if age > 15 && license {
		fmt.Println("Puede seguir avanzando")
	}

	if age <= 15 && !license {
		fmt.Println("No puede seguir circulando")
	}
}
