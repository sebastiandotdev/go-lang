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

func Params(name string, lastname string) {
	fmt.Println("!Hola, %s %s!\n", name, lastname)
}
