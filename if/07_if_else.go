package main

import "fmt"

/*Condicionales*/

/*
   if <definicion de variable> ; <condicion> {

   }
*/

/* func main() {
   	if valor := rand.int(); valor % 2 == 0 {
   		fmt.Println("El numero, ", valor, "es par")
   	} else {
   		fmt.Println("el numero, ", valor, "es impar")
   	}

   	fmt.Println("Adios !")
   }
*/

func main() {
	num := 20
	name := "sebastian"
	if num == 20 {
		fmt.Println("El numero es igual a 20")
	} else {
		fmt.Println("El numero no es igual a 20")
	}

	if name != "sebastian" {
		fmt.Println("no eres sebastian")
	} else {
		fmt.Println("bienvenido sebastian")
	}

}
