package main

import "fmt"

type person struct {
	firstName string
	lastName  string
}

func main() {
	var alex person

	alex.firstName = "Adetia"
	alex.lastName = "Marhadi"

	fmt.Println(alex)
	fmt.Printf("%+v", alex)
}
