package main

// Excercises from https://www.youtube.com/watch?v=SqrbIlUwR0U

import "fmt"

func main() {
	//Using var keyword
	var name = "Anjali"
	var age = 37
	const isCool = true
	fmt.Println(name, age, isCool)
	fmt.Printf("%T\n", name)
	fmt.Printf("%T\n", age)
	fmt.Printf("%T\n", isCool)

	//Shorthand - valid use only inside a function
	name1 := "Test"
	fmt.Println(name1, age, isCool)

	name2, email := "test", "test123@gmai.com"
	fmt.Println(name2, email, isCool)

}
