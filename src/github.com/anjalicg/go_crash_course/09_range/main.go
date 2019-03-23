package main

// Excercises from https://www.youtube.com/watch?v=SqrbIlUwR0U
import "fmt" //should be within double quote

func main() {
	ids := []int{32, 90, 45, 67, 81, 23}
	//Loops through ids
	for i, id := range ids {
		fmt.Printf("%d - ID: %d\n", i, id)
	}
	ages := map[string]int{"A": 10, "B": 20, "C": 30}
	for k, v := range ages {
		fmt.Printf("%s: %d\n", k, v)
	}

}
