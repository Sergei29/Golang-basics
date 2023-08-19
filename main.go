package main

import "fmt"

type ContactInfo struct {
	email   string
	zipCode int
}
type Person struct {
	firstName string
	lastName  string
	contact   ContactInfo
}

func (p Person) print() {
	fmt.Printf("%+v\n", p)
}

func (pointerToPerson *Person) changeContact(c ContactInfo) {
	(*pointerToPerson).contact = c
}

func (pointerToPerson *Person) changeName(newName string) {
	(*pointerToPerson).firstName = newName
}

func main() {
	jim := Person{
		firstName: "Jim",
		lastName:  "Hawkins",
		contact: ContactInfo{
			email:   "jim.admiral-benbow@gmail.com",
			zipCode: 12345,
		},
	}
	jim.print()
	jimPointer := &jim
	jimPointer.changeName("Jamezzz")
	jimPointer.changeContact(ContactInfo{email: "james@mail.com", zipCode: 99008})
	jim.print()
}
