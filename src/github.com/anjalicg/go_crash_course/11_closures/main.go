package main

// Excercises from https://www.youtube.com/watch?v=SqrbIlUwR0U
import "fmt" //should be within double quote
func adder() func(int) int {
	sum := 0
	return func(x int) int {
		sum += x
		return sum
	}
}
func main() {

	sum := adder()
	for i := 0; i < 10; i++ {
		fmt.Println(sum(i))
	}
}
