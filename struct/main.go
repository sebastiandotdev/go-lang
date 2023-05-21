package main

import "fmt"

func main() {
	type Course struct {
		Name string
		Date string
	}

	db := Course{
		Name: "err",
		Date: "09/09q&/w09e",
	}
	assets := &db

	fmt.Println(db)

	fmt.Println(assets)
}
