/*package main
import "fmt"
func main() {


	fmt.Print("Hello World")
	fmt.Println("Hello Master Academy")
	fmt.Println("Git commit testing") 
	fmt.Print("Todays Outcome:git  stages,add,push,commit ")
	fmt.Println("First text modifyed then add to stages area")
	fmt.Println("commit:stages to local repo ") 
	fmt.Println("push:local to remote repo ") 
	
}
*/
package main

import (
	"fmt"
	"strconv"
)

func main() {

	str1 := "1234"

	/** converting the str1 variable into an int using Atoi method */
	i1, err := strconv.Atoi(str1)
	if err == nil {
		fmt.Println(i1)
	}

	str2 := "5678"
	/** converting the str2 variable into an int using ParseInt method */
	i2, err := strconv.ParseInt(str2, 10, 64)
	if err == nil {
		fmt.Println(i2)
	}
}