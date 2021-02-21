package main
import "fmt"
//method
type Doctor struct {
	name string
	education string
	age int
	salary float32
}
func (d Doctor)getName() string {
	return d.name
}
func (d Doctor)getEducation() string {
	return d.education
}
func (d Doctor)getAge() int {
	return d.age
}
func (d Doctor)getSalary() float32 {
	return d.salary
}

func (d Doctor)Speak() string {
  return " can speak"
}
func main() {
	var d Doctor
	d.name = "Arafat"
	d.education = "CMC"
	d.age =24
	d.salary =500000
	fmt.Println(d)
	fmt.Println(d.Speak())
	fmt.Println(d.getName(),d.getEducation(),d.getAge(),d.getSalary())
}