package main
import "fmt"
func main() {

// x["key"] = 10
// fmt.Println(x)this get panic error without using make
x := make(map[string]int)
x["key"] = 10
x["room"]=100
x["id"]=54
fmt.Println(x["key"])
fmt.Println(x)
//delete 
delete(x, "key")
fmt.Println(x)
elements := make(map[string]string)
elements["H"] = "Hydrogen"
elements["He"] = "Helium"
elements["Li"] = "Lithium"
elements["Be"] = "Beryllium"
elements["B"] = "Boron"
elements["C"] = "Carbon"
elements["N"] = "Nitrogen"
elements["O"] = "Oxygen"
elements["F"] = "Fluorine"
elements["Ne"] = "Neon"

fmt.Println(elements["Li"])
fmt.Println(elements)
fmt.Println(elements["p"])
if name, ok := elements["Un"]; ok {
	fmt.Println(name, ok)
  }

}