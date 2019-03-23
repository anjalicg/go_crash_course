package main

// Excercises from https://www.youtube.com/watch?v=SqrbIlUwR0U

import "fmt" //should be within double quote
func greeting(name string) string {
	return "Hello " + name
}
func getSum(num1, num2 int) int { //since they are both int
	return num1 + num2
}
func main() {
	fmt.Println(greeting("World"))
	fmt.Println(getSum(5, 6))

}
