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
	// var book1 book
	// book1.Title = "An Introduction to Programming in Go"
	// book1.Author = "Caleb Doxsey"
	// book1.ISBN = "978-1478355823"
	// book1.Price =99.99
	// book1.Pages=165
	// fmt.Println(book1)
	// fmt.Println(book1.ISBN)

	//way 2 annonymous struct
	book1:=struct{
		Title string
	Author string
	ISBN string
	Price float64
	Pages int
	}{
		Title : "An Introduction to Programming in Go",
		Author : "Caleb Doxsey",
		ISBN : "978-1478355823",
		Price : 99.99,
		Pages: 165,

	}
	fmt.Println(book1)
	fmt.Println(book1.ISBN)

	//way 3 
	book2 := book{
		Title : "Introducing Go",
		Author : "Caleb Doxsey",
		ISBN : "9781491941959",
		Price : 13.75,
		Pages: 165,
	}
	fmt.Println(book2)
	fmt.Println(book1.ISBN)
}