package main
import "fmt"
type Email interface {
	write(string,string,string,string)
	send()
	read() 
}

type Test struct {
	To string
	From string
	Subject string
	Message string
}
func (e Test)send(){
	fmt.Println("email send to ", e.To," succesfully")
   }
func (e Test)read(){
		fmt.Println("receievd message from ", e.From," succesfully")
   }
func (e Test)write(to string, from string,subject string,message string){
 fmt.Println("To:",to," from:" ,from,"subject:",subject,"Body:",message)
}

func main() {
var t Test
t.To = "nahidulss@gmail.com"
t.From = "u1504054@student.cuet.ac.bd"
t.Subject = "test msg"
t.Message="master acdemy Bangladesh"
t.write(t.To,t.From,t.Subject,t.Message)
t.send()
t.read()
}