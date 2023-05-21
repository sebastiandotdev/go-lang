package footer

import "fmt"

func arrays() {
	var food [3]string
	food[0] = "mzn"
	food[1] = "prz"
	food[2] = "cco"

	food_two := [3]string{"hsz", "jsjs", "su"}
	fmt.Println(food)
	fmt.Println(food_two)
}
