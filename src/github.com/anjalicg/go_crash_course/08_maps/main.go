package main

// Excercises from https://www.youtube.com/watch?v=SqrbIlUwR0U
import "fmt"

//should be within double quote

func main() {
	//Define a map
	emails := make(map[string]string)
	//Assign key values
	emails["Bob"] = "bob@gmail.com"
	emails["Blah"] = "blah@gmail.com"
	emails["dummy"] = "dummy@gmail.com"
	fmt.Println(emails)
	fmt.Println(len(emails))
	fmt.Println(emails["Bob"])

	//delete
	delete(emails, "Bob")
	fmt.Println(emails)

	//Declare map and assign key value
	ages := map[string]int{"A": 10, "B": 20, "C": 30}
	fmt.Println(ages)
	fmt.Println(len(ages))
	fmt.Println(ages["B"])
	delete(ages, "B")
	fmt.Println(ages)
	delete(ages, "B")
	fmt.Println(ages)
}
