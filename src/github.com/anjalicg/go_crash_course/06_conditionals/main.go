package main

// Excercises from https://www.youtube.com/watch?v=SqrbIlUwR0U
import "fmt" //should be within double quote

func main() {

	x := 15
	y := 15
	if x <= y {
		fmt.Printf("%d is less than or equal to %d\n", x, y)
	} else {
		fmt.Printf("%d is more than %d\n", x, y)
	}

	if x < y {
		fmt.Printf("%d is less than  %d\n", x, y)
	} else if x > y {
		fmt.Printf("%d is more than %d\n", x, y)
	} else {
		fmt.Printf("%d is equal to  %d\n", x, y)
	}

	//Switch
	switch x {
	case 15:
		fmt.Println("It is 15")
	default:
		fmt.Println("It is not 15")

	}

}
