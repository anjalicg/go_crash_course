package main

// Excercises from https://www.youtube.com/watch?v=SqrbIlUwR0U
import "fmt" //should be within double quote

func main() {
	var fruitArr [2]string
	fruitArr[0] = "Apple"
	fruitArr[1] = "Orange"
	fmt.Println(fruitArr)
	// Declare and assign at the same time
	fruitArr1 := [2]string{"Apple", "Orange"}
	fmt.Println(fruitArr1)
	//Slices
	fruitSlice := []string{"A", "B", "C", "D"}
	fmt.Println(fruitSlice)
	fmt.Println(len(fruitSlice))
	fmt.Println(fruitSlice[1:3])
}
