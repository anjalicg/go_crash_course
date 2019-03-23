package main

// Excercises from https://www.youtube.com/watch?v=SqrbIlUwR0U
import (
	"fmt"
	"strconv"
)

//Define Person struct
type Person struct {
	// firstName string
	// lastName  string
	// city      string
	// gender    string
	firstName, lastName, city, gender string
	age                               int
}

// Greeting method - value receiver
func (p Person) greet() string {
	return "Hello my name is " + p.firstName + " " + p.lastName + " and i am " + strconv.Itoa(p.age) + "\n"
}

// has birthday method - pointer receiver
func (p *Person) hasBirthday() {
	p.age++
}

// getMarried - pointer receiver
func (p *Person) getMarried(spouseLastName string) {
	if p.gender == "M" {
		return
	} else {
		p.lastName = spouseLastName
	}

}

func main() {

	person1 := Person{firstName: "Anjali", lastName: "CG", city: "SC", gender: "F", age: 37}
	fmt.Println(person1)

	person2 := Person{"Anjali", "CG", "SC", "F", 37}
	fmt.Println(person2)
	fmt.Println(person2.firstName)
	person2.age++
	fmt.Println(person2.age)
	person2.hasBirthday()
	fmt.Println(person2.greet())
	person2.getMarried("Suresh")
	fmt.Println(person2.greet())

	person3 := Person{firstName: "Suresh", lastName: "G", city: "SC", gender: "M", age: 37}
	person3.getMarried("CG")
	fmt.Println(person3.greet())

}
