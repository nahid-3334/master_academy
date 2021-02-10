package main
import "fmt"
func main()  {
	var name string
	var age int
	fmt.Println("Enter your name & Age:")
	fmt.Scanf("%s %d",&name,&age) 
	fmt.Println("name:",name," Age:",age)
}