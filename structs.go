package main

import "fmt"

func main() {
	var person1 person

	person1.name = "Kwaleyela Ikafa"
	person1.age = 21
	person1.gender = "Male"
	person1.address = "Woodlands ext"

	getPersonDetails(person1)
	fmt.Printf("\n")
	setNewAddress(&person1, "Winela, Nalukota")
	getPersonDetails(person1)
}

type person struct {
	name    string
	age     int
	gender  string
	address string
}

func getPersonDetails(p person) {
	fmt.Println("Name:", p.name)
	fmt.Println("Age:", p.age)
	fmt.Println("Gender:", p.gender)
	fmt.Println("Address:", p.address)
}

func setNewAddress(p *person, address string) {
	p.address = address
}
