//devhints cheat sheat
package main

import "fmt"
func main() {
	
	//Number 
	fmt.Println("Number type:")
	num1 := 3          // int
num2 := 3.5       // float64
num3 := 3 + 4i     // complex128
num4 := byte('a')  // byte (alias for uint8)
fmt.Println(num1, num2, num3, num4)
var myHexNumber = 0xFF  // Use prefix '0x' or '0X' for declaring hexadecimal numbers
	var myOctalNumber = 034 // Use prefix '0' for declaring octal numbers
	//%#x, %#o\n f
	fmt.Println(myOctalNumber,myHexNumber)
	//char ->byte ,rune
	var myByte byte = 'a'
	var myRune rune = '♥'

	fmt.Printf("%c = %d and %c = %U\n", myByte, myByte, myRune, myRune)
	//string
	fmt.Println("Number type:")
	name := "\nnahid\t"//normal string
	namemulti:=`I am Nahid
	CUET CSE-15
	`//back tick for multi line raw string
	fmt.Println(name)
	
	fmt.Println(namemulti)
	//convertion
 marks1 :=75
 marks2 :=25.50
 var sum=float64(marks1) +marks2 
 fmt.Println(sum)
}
