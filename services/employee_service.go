package services

import (
	"employee-managment/models"
)

type EmployeeService struct {
	Employees map[string]*models.Employee
}

func NewEmployeeService() *EmployeeService {
	return &EmployeeService{Employees: make(map[string]*models.Employee)}
}

func (s *EmployeeService) AddEmployee(employee *models.Employee) {
	s.Employees[employee.ID] = employee
}

func (s *EmployeeService) GetEmployeeByID(id string) (*models.Employee, bool) {
	employee, exists := s.Employees[id]
	return employee, exists
}

func (s *EmployeeService) GetAllEmployees() []*models.Employee {
	var employees []*models.Employee
	for _, employee := range s.Employees {
		employees = append(employees, employee)
	}
	return employees
}
