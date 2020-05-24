package main

import "fmt"

type person struct {
	firstName string
	lastName  string
	contactInfo
}

type contactInfo struct {
	email   string
	zipCode int
}

func main() {
	// p1 := person{"Avinash", "Kunuje"}
	// p := person{firstName: "Avinash", lastName: "Kunuje"}
	// fmt.Println(p)
	// var p person
	// p.firstName = "Avinash"
	// p.lastName = "KS"
	// fmt.Println(p)
	// fmt.Printf("%+v", p)

	p := person{
		firstName: "Avinash",
		lastName:  "Kunuje",
		contactInfo: contactInfo{
			email:   "itsavi.ks@gmail.com",
			zipCode: 560102,
		},
	}
	fmt.Println("ggggg", &p)
	p.updateName("Kavya")
	p.print()
}
func (pointerToPerson *person) updateName(newFirstName string) {
	(*pointerToPerson).firstName = newFirstName
}

func (pointerToPerson person) print() {
	fmt.Printf("%+v", pointerToPerson)
}
