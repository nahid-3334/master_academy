package main 
import "fmt"
type book struct {
	Title string
	Author string
	ISBN string
	Price float64
	Pages int
}
func main() {
	//struct 
	var book1 book
	book1.Title = "An Introduction to Programming in Go"
	book1.Author = "Caleb Doxsey"
	book1.ISBN = "978-1478355823"
	book1.Price =99.99
	book1.Pages=165
	fmt.Println(book1)
	fmt.Println(book1.ISBN)
}