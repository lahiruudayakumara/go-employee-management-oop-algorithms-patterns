package services

import "employee-managment/models"

type DepartmentService struct {
	Departments map[string]*models.Department
}

func NewDepartmentService() *DepartmentService {
	return &DepartmentService{Departments: make(map[string]*models.Department)}
}

func (s *DepartmentService) AddDepartment(department *models.Department) {
	s.Departments[department.ID] = department
}

func (s *DepartmentService) GetDepartmentByID(id string) (*models.Department, bool) {
	department, exists := s.Departments[id]
	return department, exists
}
