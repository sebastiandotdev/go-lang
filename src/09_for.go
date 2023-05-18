package src

import "fmt"

/*
	for {
		// intrucciones a repetir
	}
*/

func main() {
	var a = 3
	// for {
	// 	fmt.Print("Salir? (s/n): ")
	// 	var c rune
	// 	fmt.Scanf("%c\n", &c)
	// 	if c == 'N' || c == 'n' {
	// 		continue
	// 	}
	// 	if c == 'S' || c == 's' {
	// 		break
	// 	}
	// }
	// fmt.Println("Adios !")
	for_go()
	if a == 3 {
		i := 1
		fmt.Println("'i' solo es visible en este contexto: ", i)
	}
	fmt.Println("'a' solo es visible en todo main: ", a)
	b := 0
	if true {
		a := 1
		b = 1
		a++
		b++
	}
	fmt.Println("a = %d b = %d\n", a, b)
}

func for_go() {
	for i := 1; i <= 10; i++ {
		fmt.Println(i)
	}
}
