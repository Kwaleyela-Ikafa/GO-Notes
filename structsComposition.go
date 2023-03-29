// Structs with composition
// Structs dont have inheritance they use composition

package main

import "fmt"

func main() {
	con1 := contact{
		"Kwaleyela",
		"Ikafa",
		"+260974036274",
	}

	bus1 := business{
		"WesTTesch",
		"Woodlands",
		con1,
	}

	bus1.info()
}

type contact struct {
	fName string
	lName string
	phone string
}

type business struct {
	name    string
	address string
	contact
}

func (b business) info() {
	fmt.Printf(
		"Contact at %s is %s %s \n",
		b.name, b.contact.fName,
		b.contact.lName,
	)
}
