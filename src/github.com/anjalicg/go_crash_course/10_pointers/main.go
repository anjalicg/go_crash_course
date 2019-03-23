package main

// Excercises from https://www.youtube.com/watch?v=SqrbIlUwR0U
import "fmt" //should be within double quote

func main() {
	a := 5
	b := &a
	fmt.Println(a, b, *b) // Use * to read value from address
	fmt.Printf("%T\n", a)
	fmt.Printf("%T\n", b)
	*b = 20
	fmt.Printf("%d\n", a)
}
