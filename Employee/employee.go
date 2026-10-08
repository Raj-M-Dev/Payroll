package main

type Employee struct {
	name   string
	age    int
	gender string
	role   Role
}

type Role struct {
	id   string
	name string
}
