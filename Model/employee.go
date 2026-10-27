package Model

type Employee struct {
	Name   string
	Age    int
	Gender string
	Role   Role
}

type Role struct {
	Id   string
	Name string
}
