package main

import "fmt"

func main() {

	mapa := map[string]string{
		"10": "notebook",
		"15": "tv",
		"21": "heladera",
		"27": "monitor",
		"35": "camara",
	}

	array := []string{}

	for {

		var valueNew string
		fmt.Scanln(&valueNew)

		if valueNew == "0" {
			fmt.Println("Los valoresd del array son: ", array)

			break
		}

		if item, ok := mapa[valueNew]; ok {
			array = append(array, item)
		} else {
			array = append(array, "No encontrado")
		}
	}
}
