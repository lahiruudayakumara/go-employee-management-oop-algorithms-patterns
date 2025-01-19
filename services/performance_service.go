package services

import "employee-managment/models"

type PerformanceService struct {
	Performances map[string]*models.Performance
}

func NewPerformanceService() *PerformanceService {
	return &PerformanceService{Performances: make(map[string]*models.Performance)}
}

func (s *PerformanceService) AddPerformance(performance *models.Performance) {
	s.Performances[performance.EmployeeID] = performance
}

func (s *PerformanceService) GetPerformanceByEmployeeID(employeeID string) (*models.Performance, bool) {
	performance, exists := s.Performances[employeeID]
	return performance, exists
}
