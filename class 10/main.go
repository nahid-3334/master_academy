package main

import "fmt"
func main()  {
	//Array
	var students [5]string
	fmt.Println(students)
	fmt.Println(len(students))
	//len for length empty string for default
	//set 
	students[0] = "Apu"
	students[1]="Robin"
	students[2]="John"
	students[3]="Robiul"
	students[4]="Foysal"
	//get
	fmt.Println(students[3])
	fmt.Println(students[4])
	//short hand implicit=...
	marks:=[...]int{80,40,70,87,76}
	fmt.Println(marks)
}