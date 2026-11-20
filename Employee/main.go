package main

import (
	"fmt"
	"net/http"
	model "paySlip/model"

	"github.com/gin-gonic/gin"
)

type EmployeIN interface {
	addEmployee() error
	//	getEmployee() error
	//	editEmployee() error
	//	deleteEmployee() error
}

func hander() {
	server := gin.Default()
	//server.GET("/employees", getEmployes)
	server.GET("/test", getTest)
	server.Run(":8080")
}
func main() {
	hander()
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

// func getEmployes(c *gin.Context, e *model.Employee) {

// 	employees := []model.Employee{}
// 	e.Name = "Suresh"
// 	e.Age = 34
// 	e.Gender = "Male"
// 	e.Role = model.Role{Id: "002", Name: "HR"}
// 	employees = append(employees, *e)
// 	c.JSON(http.StatusOK, employees)

// }

func getTest(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"message": "Hello World",
	})
}
