package src

import "fmt"

/*
func <Nombre de la funcion>() {
	// codigo a ejecutar en cada invocacion
}

*/

func Hola() {
	fmt.Println("Hola!")
}

func Params(name, lastname string) {
	fmt.Println("!Hola, %s %s!\n", name, lastname)
}

func Max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func MaxMin(a, b int) (int, int) {
	if a > b {
		return a, b
	}
	return b, a
}

func MinMax(a, b int) (max, min int) {
	if a > b {
		max = a
		min = b
	} else {
		min = a
		max = b
	}
	return
}

// func() {
// 	fmt.Println("hello")
// }()
