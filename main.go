package main

import (
	"Employee/emp"
	"fmt"
)

func main() {

	e := emp.Employee{Name: "Mahesh", Age: 30, Gender: "Male", Salary: 50000.00}
	fmt.Println(e)
}
