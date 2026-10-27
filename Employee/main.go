package main

import (
	"fmt"
	model "paySlip/model"
)

type EmployeIN interface {
	addEmployee() error
	//	getEmployee() error
	//	editEmployee() error
	//	deleteEmployee() error
}

func main() {
	emp := model.Employee{}
	err := addEmployee(&emp)
	if err != nil {
		fmt.Println("Error adding employee:", err)
	}
}
func addEmployee(e *model.Employee) error {

	employees := []model.Employee{}
	e.Name = "Mahesh"
	e.Age = 30
	e.Gender = "Male"
	e.Role = model.Role{Id: "001", Name: "Manager"}
	employees = append(employees, *e)
	fmt.Println(employees)
	return nil
}
