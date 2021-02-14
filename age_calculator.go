package main

import "fmt"
import "time"
import "strconv"
func main() {
	var name string="nahid"
	// fmt.Print("Give your Name:")
	// fmt.Scanf("%s", &name)
	current_time :=time.Now();
	var birthday string  
	fmt.Print("Give Your birthday as year-month-date format(2000-01-15) :")
	fmt.Scanf("%s", &birthday)
	birthyear :=birthday[:4]
	birthmonth :=birthday[5:7]
	birthdate:=birthday[8:10]

	current:=current_time.Format("2006-01-02")
	currentdate:=current[8:10]
	currentmonth:=current[5:7]
	currentyear:=current[:4]
	//string to int
	birth_year, err:= strconv.Atoi(birthyear)
	birth_month,err := strconv.Atoi(birthmonth)
	birth_date, err := strconv.Atoi(birthdate)
	if err == nil {
		fmt.Println("birth year:",birth_year," month:",birth_month," date:",birth_date)
	}
	//string to int
	current_year, err:= strconv.Atoi(currentyear)
	current_month,err := strconv.Atoi(currentmonth)
	current_date, err := strconv.Atoi(currentdate)
	if err == nil {
		//fmt.Println(birth_year,birth_month,birth_date)
		fmt.Println("current year:",current_year, " month:",current_month," date:",current_date)
	}
	//fmt.Println(birth_year,birth_month,birth_date)
	
	 month:= [12]int{ 31, 28, 31, 30, 31, 30, 31,  
		31, 30, 31, 30, 31 }

// if birth date is greater then current birth 
// month then do not count this month and add 30  
// to the date so as to subtract the date and 
// get the remaining days 
if birth_date > current_date { 
current_date = current_date + month[birth_month - 1]
current_month = current_month - 1
} 

// if birth month exceeds current month, then do 
// not count this year and add 12 to the month so 
// that we can subtract and find out the difference 
if (birth_month > current_month) { 
current_year = current_year - 1
current_month = current_month + 12 
}

// // calculate date, month, year 
calculated_date :=current_date- birth_date
calculated_month := current_month - birth_month 
calculated_year  := current_year - birth_year
fmt.Printf("Mr. %s current",name)
fmt.Println(" Age:",calculated_year," Years ", calculated_month , " Month ",calculated_date ," Days")
}