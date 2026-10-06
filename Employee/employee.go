package Employee

type employee struct {
	Name String
	Age int
	Gender string
	Salary float64
	Role role
}

type role struct {
	id string
	name string
}