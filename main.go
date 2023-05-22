package main

import "fmt"

func main() {
	/*gitpod cloud*/
	a := 3
	fmt.Println("antes: ", a)
	Increment(&a)
	fmt.Println("despues: ", a)
	/*
		Go
	*/
	fmt.Println("Hola! Go")
}

func Increment(value *int) {
	*value++
}
