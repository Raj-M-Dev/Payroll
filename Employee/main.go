package main

import (
	"fmt"
)

func main() {
	var emp Employee
	employees := []Employee{}
	emp = Employee{name: "Mahesh", age: 30, gender: "Male", role: Role{id: "1", name: "Developer"}}
	employees = append(employees, emp)
	fmt.Println(employees)

}
