package main

import "fmt"

type person struct {
	name    string
	age     int
	favFood []string
}

func main() {
	// favFood := []string{"kimchi", "ramen"}
	// nico := person{"박충호", 29, favFood}
	// fmt.Println(nico)

	favFood := []string{"kimchi", "ramen"}
	nico := person{name: "박충호", age: 29, favFood: favFood}
	fmt.Println(nico)
}
