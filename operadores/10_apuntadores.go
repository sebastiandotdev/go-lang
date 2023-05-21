package main

import "fmt"

var pi *int

func main() {
	pi = nil
	i := 10
	p := &i

	a := *p
	*p = 21

	fmt.Println(a)
	fruit := "manzana"

	fmt.Printf("Tipo de dato: %T, valor: %s", fruit, fruit)

	if pi == nil {
		fmt.Println("No puedo hacer nada con este apuntador")
		fmt.Println("Porque no apunta a nada!")
	}
}
