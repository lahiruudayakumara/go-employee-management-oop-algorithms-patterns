package patterns

import "employee-managment/models"

type SalaryCalculator interface {
	CalculateSalary(employee *models.Employee) float64
}

type FullTimeSalary struct{}

func (s *FullTimeSalary) CalculateSalary(employee *models.Employee) float64 {
	return employee.Salary
}

type PartTimeSalary struct{}

func (s *PartTimeSalary) CalculateSalary(employee *models.Employee) float64 {
	return employee.Salary / 2
}

type ContractSalary struct{}

func (s *ContractSalary) CalculateSalary(employee *models.Employee) float64 {
	return employee.Salary / 3
}
