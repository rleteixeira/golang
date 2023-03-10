package main

import "fmt"

type contactInfo struct {
	email string
	zipCode int
}

type person struct {
	firstName string
	lastName  string
	contact contactInfo			//can use just the type without name
}

func main() {
	/*
	alex := person{firstName: "Alex", lastName: "Anderson"}
	alex := person{"Alex", "Anderson"}
	fmt.Println(alex)
	*/	

	/*
	var alex person				//empty string
	fmt.Println(alex)			//{ }
	fmt.Printf("%+v", alex)		//{firstName: lastName:}
	*/

	/*
	var alex person	
	alex.firstName = "Alex"	
	alex.lastName = "Anderson"		
	fmt.Println(alex)			
	fmt.Printf("%+v", alex)		
	*/

	jim := person {
		firstName: "Jim",
		lastName: "Party",
		contact: contactInfo{
			email: "jim@gmail.com",
			zipCode: 94000,
		},
	}
	
	/*
	jimPointer := &jim
	jimPointer.updateName("jimmy")
	jim.print()
	*/

	jim.updateName("jimmy")
	jim.print()
}

func (pointerToPerson *person) updateName(newFirstName string) {
	(*pointerToPerson).firstName = newFirstName
}

func (p person) print(){
	fmt.Printf("%+v", p)
}