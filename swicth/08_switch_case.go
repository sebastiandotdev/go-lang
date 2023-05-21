package src

import (
	"fmt"
	"runtime"
)

/* switch <expresion a evaluar> {
	case <valor posible>:
		// codigo
	case <valor posible>:
	// CODIGO
	default:
		// codigo a ejecutar si la expresion no toma ninguna de los posibles valores anteriores
} */

/*
   switch <inicializacion de variable>; <expresion a evaluar>
*/

func main() {
	arch := runtime.GOARCH

	switch arch {
	case "386":
		fmt.Println("x86 de 32bits")
	case "amd64":
		fmt.Println("x86 de 64bits")
	default:
		fmt.Println("error")
	}
}
