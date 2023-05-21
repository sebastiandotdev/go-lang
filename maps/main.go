package main

import "fmt"

func main() {
	animals := make(map[string]string)

	animals["pets"] = "perro"

	fmt.Println(animals)

	fruits := map[string]string{
		"apple": "manzana",
	}
	fmt.Println(fruits)
}
