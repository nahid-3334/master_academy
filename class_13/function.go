package main
import "fmt"
//Function with Parameters:
func add(x int,y int){
	r:=x+y
	fmt.Println("Sum",r)
	}
	//Function with Return Type:
func sub(x int,y int) int {
		r:=x-y
		return r
		}	
//Named Return Values		
func mul(x int,y int) (r int) {
			r=x*y
			return 
			}	
func div(x int,y int) int {
				r:=y/x
				return r
				}	
//Returning Multiple Values
func circle(r float64) (area , parameter float64) {
	parameter = 2 *3.1416*r
	area = 3.1416 * r* r
	return // Return statement without specify variable name
}	
//Passing Address to a Function
func update(a *int, t *string) {
	*a = *a + 5      // derefrencing pointer address
	*t = *t + " Islam" // derefrencing pointer address
	return
}	
// area: = func(l int, b int) int {
// 	return l * b
// }					
func main() {
 add(10,20)
 fmt.Println("Subtract:",sub(10,20))
 fmt.Println("MUltiplication:",mul(10,20))
 fmt.Println("Division:",div(10,20))
 var a, p float64
	a, p = circle(5.5)
	fmt.Println("Area:", a)
	fmt.Println("Parameter:", p)
	//pointer 
	var age = 18
	var name = "Nahid"
	fmt.Println("Before:", name, age,"years old")

	update(&age, &name)

	fmt.Println("After :", name, age,"years old")
	area:= func(l int, b int) int {
			return l * b
		}
	fmt.Println(area(20, 30))
}	