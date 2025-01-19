package models

type Employee struct {
	ID         string
	Name       string
	Age        int
	Department string
	Salary     float64
}

func NewEmployee(id, name, department string, age int, salary float64) *Employee {
	return &Employee{
		ID:         id,
		Name:       name,
		Age:        age,
		Department: department,
		Salary:     salary,
	}
}
