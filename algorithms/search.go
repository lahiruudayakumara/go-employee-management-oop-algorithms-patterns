package algorithms

import "employee-managment/models"

// LinearSearch searches for an employee by name in the data slice.
func LinearSearch(data []*models.Employee, query string) *models.Employee {
	for _, emp := range data {
		if emp.Name == query {
			return emp
		}
	}
	return nil
}
