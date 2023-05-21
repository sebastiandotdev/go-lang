package main

import (
	"fmt"
	"io/ioutil"
)

func main() {
	file, err := ioutil.ReadFile("./somethings.txt")

	if err != nil {
		panic(err)
	}
	fmt.Println(file)
}
