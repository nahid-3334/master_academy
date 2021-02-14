package main
import "fmt"
func main() {
	//flow control
	//age grouping if else
	// var age int
	// fmt.Println("Enter yor age:")
	// fmt.Scanf("%d", &age)
	// if age < 3{
	// 	fmt.Println("infant")
	// }else if age >=3 && age <=12{
	// 	fmt.Println("Child")
	// }else if age>=13 && age <=18{
	// 	fmt.Println("Teenage")
	// }else{
	// 	fmt.Println("adult")
	// }
	//marks group using switch
	//fixed value
	// var marks int
	// fmt.Println("Enter yor marks:")
	// fmt.Scanf("%d", &marks)
// >=80	A+
// 70-79	A
// 60-69	B
// 50-59   C
// 40-49	D
	// switch marks/10 {
	// case 4:
	// 	fmt.Println("your grade D")
	// case 5:
	// 	fmt.Println("your grade C")
	// case 6:
	// 	fmt.Println("your grade B")
	// case 7:
	// 	fmt.Println("your grade A")
	// case 8,9,10:
	// 	fmt.Println("your grade A+")

	// }
	//loop 
	//mormal
	for i := 0; i<10; i++ {
		fmt.Println(i,"nahid")
	}
	//get value from array range loop
 	students :=[]string{"apu","foysal","Robiul"}
	for i,std:=range students{
		fmt.Println(i,std)
	}
	//while in for
	i:=0
	for i<10{
		fmt.Println(i,"nahid")
		i++
	}
}