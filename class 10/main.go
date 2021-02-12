package main

import "fmt"
import "reflect"
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
	//slice =sub array 
	//short way & from array  hold starting index to last value-1 
	twostudents:=students[1:2]
	fmt.Println(twostudents)
	//direct way by makes
	x:=make([]string,5,10)
	fmt.Println(x)
	 var fruits []string
	 fruits=append(fruits,"apple","mango","orange","banana")
	 //Type
	 fmt.Printf("%T %s\n",fruits,fruits)
	 //type by reflect
	 a:=reflect.TypeOf(fruits).Kind().String()
	 fmt.Println(a)
	 //Here is an example of copy:
	slice1 := []int{1,2,3}
	slice2 := make([]int, 2)
	copy(slice2, slice1)
	fmt.Println(slice1, slice2)
}