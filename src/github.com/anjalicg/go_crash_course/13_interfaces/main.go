package main

// Excercises from https://www.youtube.com/watch?v=SqrbIlUwR0U
import (
	"fmt"
	"math"
)

type Shape interface {
	area() float64
}
type Circle struct {
	x, y, radius float64
}
type Rectangle struct {
	length, width float64
}

func (c Circle) area() float64 {
	return math.Pi * c.radius * c.radius
}

func (r Rectangle) area() float64 {
	return r.width * r.length
}

func getArea(s Shape) float64 {
	return s.area()
}
func main() {
	circle := Circle{0, 0, 5}
	rectangle := Rectangle{10, 20}
	fmt.Printf("Circle Area: %f\n", getArea(circle))
	fmt.Printf("Rectangle Area: %f\n", getArea(rectangle))

}
