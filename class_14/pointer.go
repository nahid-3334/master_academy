package main
import "fmt"
func update(s *int){
	fmt.Println("s=",*s)
	*s= *s +10
}
func main() {
	var x int
	var y *int
	fmt.Println("zero value x: ", x,"\nzero value y",y)
	fmt.Println("address x: ",&x," address y:",&y)
	x=10//assign
	y= &x //referencing
	fmt.Println("value x: ", x,"\nvalue y",y)//accessing
	fmt.Println("y dererencing value is", *y)//dereferencing
	*y=20
	fmt.Println("value x: ", x,"\ndereference value y:",*y)
	update(y)
	fmt.Println("value x=",x)
}