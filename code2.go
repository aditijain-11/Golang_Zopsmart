package main

import "fmt"

func main() {
	fmt.Println("Hello ")
	var name string
	name = "Sakshi"

	fmt.Printf("Welcome to %s\n", name)

	var (
		i int     = 13
		f float64 = 5.14
	)

	fmt.Printf("Type of i: %T\n", i)
	fmt.Printf("Type of f: %T\n", f)
}
