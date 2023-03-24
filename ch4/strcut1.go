package main

import (
	"fmt"
	"time"
)

type Employee struct {
	ID            int
	Name, Address string
	DoB           time.Time
	Position      string
	Salary        int
	ManagerId     int
}

var employee = Employee{Position: "manager"}

func main() {

	EmployeeById(1).Salary = 1000
	fmt.Println(employee)
}

func EmployeeById(id int) *Employee {

	if id == 1 {
		return &employee
	}
	return &Employee{}
}
