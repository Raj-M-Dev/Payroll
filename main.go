package main

import "fmt"

type employee struct {
	Name   string
	Age    int
	Salary float64
}

func main() {
	e := employee{Name: "Mahesh", Age: 30, Salary: 50000.00}
	fmt.Println(e)
}
