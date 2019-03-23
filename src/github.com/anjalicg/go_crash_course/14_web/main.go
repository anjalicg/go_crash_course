package main

// Excercises from https://www.youtube.com/watch?v=SqrbIlUwR0U
import (
	"fmt"
	"net/http"
)

func index(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "<h1>Hello World</h1>")
}
func main() {
	http.HandleFunc("/", index)
	fmt.Println("Server starting..")
	http.ListenAndServe(":3000", nil)

}
