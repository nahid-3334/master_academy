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
// elements := make(map[string]string)
// elements["H"] = "Hydrogen"
// elements["He"] = "Helium"
// elements["Li"] = "Lithium"
// elements["Be"] = "Beryllium"
// elements["B"] = "Boron"
// elements["C"] = "Carbon"
// elements["N"] = "Nitrogen"
// elements["O"] = "Oxygen"
// elements["F"] = "Fluorine"
// elements["Ne"] = "Neon"

// fmt.Println(elements["Li"])
// fmt.Println(elements)
// fmt.Println(elements["p"])
// students[0] = "Apu"
// 	students[1]="Robin"
// 	students[2]="John"
// 	students[3]="Robiul"
// 	students[4]="Foysal"
// 	//get
// 	fmt.Println(students[3])
// 	fmt.Println(students[4])
// 	//short hand implicit=...
// 	marks:=[...]int{80,40,70,87,76}
// 	fmt.Println(marks)


//maps of maps
elements := map[string]map[string]string{
    "H": map[string]string{
      "name":"Hydrogen",
      "state":"gas",
    },
    "He": map[string]string{
      "name":"Helium",
      "state":"gas",
    },
    "Li": map[string]string{
      "name":"Lithium",
      "state":"solid",
    },
    "Be": map[string]string{
      "name":"Beryllium",
      "state":"solid",
    },
    "B":  map[string]string{
      "name":"Boron",
      "state":"solid",
    },
    "C":  map[string]string{
      "name":"Carbon",
      "state":"solid",
    },
    "N":  map[string]string{
      "name":"Nitrogen",
      "state":"gas",
    },
    "O":  map[string]string{
      "name":"Oxygen",
      "state":"gas",
    },
    "F":  map[string]string{
      "name":"Fluorine",
      "state":"gas",
    },
    "Ne":  map[string]string{
      "name":"Neon",
      "state":"gas",
    },
  }
  if el, ok := elements["UN"]; ok {
    fmt.Println(el["name"], el["state"])
  }
}