package patterns

import "employee-managment/models"

type EmployeeType string

const (
	FullTimeEmployee EmployeeType = "FullTime"
	PartTimeEmployee EmployeeType = "PartTime"
	ContractEmployee EmployeeType = "Contract"
)

func EmployeeFactory(empType EmployeeType, id, name, department string, age int, salary float64) interface{} {
	switch empType {
	case FullTimeEmployee:
		return models.NewEmployee(id, name, department, age, salary)
	case PartTimeEmployee:
		return models.NewEmployee(id, name, department, age, salary/2)
	case ContractEmployee:
		return models.NewEmployee(id, name, department, age, salary/3)
	default:
		return nil
	}
}
