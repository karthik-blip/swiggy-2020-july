package main

import (
	"fmt"
	"strconv"
)

type Address struct {
	hno int
	street string
	city string
	pin int
}

type  employee struct {
	empID int
	empName string
	salary int
	address Address
}

type animal interface {
	eat(i int) string
	breath(i int) string
	sleep(i int) string
	print_details(i int)
}

func (e employee) eat (i int)(string){
	return (strconv.Itoa(i)+" Eating")
}

func (e employee) breath(i int)(string){
	return(strconv.Itoa(i)+" Breating")
}

func (e employee) sleep(i int)(string){
	return(strconv.Itoa(i)+" sleeping")
}

func (e employee) print_details(i int)(){
	fmt.Println("EMPLOYEE " , i)
	fmt.Println("name" , e.empName)
	fmt.Println("employee id" , e.empID)
	fmt.Println("employee salary " , e.salary)
	fmt.Println("employee house number", e.address.hno)
	fmt.Println("employee street",e.address.street)
	fmt.Println("employee city",e.address.city)
	fmt.Println("employee PIN",e.address.pin)
}



func main() {
	var employees []animal
	var cond string
	employees = make([]animal, 0, 5)
	i := 0
	var emp_id, salary, hno, pin int
	var emp_name, street, city string
	for true {
		//employees = append(employees , animal(employee{0 , "0" , 0 , Address{0,"0","0",0}}))
		fmt.Println("Enter Employee ID, employee Name , salary for employee ", i+1)
		fmt.Scanln(&(emp_id))
		fmt.Scanln(&(emp_name))
		fmt.Scanln(&(salary))
		fmt.Println("Enter house number, street name, city, PIN of the employee ", i+1)
		fmt.Scanln(&(hno))
		fmt.Scanln(&(street))
		fmt.Scanln(&(city))
		fmt.Scanln(&(pin))
		fmt.Println("Do you want to quit ? (n/q)")
		employees = append(employees, animal(employee{empID: emp_id, empName: emp_name, salary: salary, address: Address{hno: hno, street: street, city: city, pin: pin}}))
		fmt.Scanln(&cond)
		if cond == "q" {
			break
		}
		i++
	}

	for i,emp := range employees{
		emp.print_details(i)
		fmt.Println()
		fmt.Println(emp.breath(i + 1))
		fmt.Println(emp.sleep(i + 1))
		fmt.Println(emp.eat(i + 1))
		fmt.Println()
	}

}