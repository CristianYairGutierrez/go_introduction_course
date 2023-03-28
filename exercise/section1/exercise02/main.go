package main

import "fmt"

func main() {

	array := []int{}

	for {

		fmt.Println("Ingresa un valor para agregar al array: ")

		var valueNew = 0
		fmt.Scanln(&valueNew)
		if valueNew == 0 {
			break
		}

		array = append(array, valueNew)
	}

	fmt.Println("Los valoresd del array son: ", array)
}
