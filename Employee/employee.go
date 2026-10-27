package main

import (
	"fmt"
)

type Employee struct {
	name   string
	age    int
	gender string
	role   Role
}
type EmployeIN interface {
	addEmployee() error
	getEmployee() error
	editEmployee() error
	deleteEmployee() error
}
type Role struct {
	id   string
	name string
}

func (e *Employee) addEmployee() error {

	employees := []Employee{}
	e.name = "Mahesh"
	e.age = 30
	e.gender = "Male"
	e.role = Role{id: "001", name: "Manager"}
	employees = append(employees, *e)
	fmt.Println(employees)
	return nil
}
