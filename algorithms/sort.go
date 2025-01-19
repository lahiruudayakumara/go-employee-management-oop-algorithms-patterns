package algorithms

import "employee-managment/models"

// SortBySalary sorts the employees by their salary in ascending order.
func SortBySalary(data []*models.Employee) []*models.Employee {
	for i := 0; i < len(data)-1; i++ {
		for j := 0; j < len(data)-i-1; j++ {
			if data[j].Salary > data[j+1].Salary {
				data[j], data[j+1] = data[j+1], data[j]
			}
		}
	}
	return data
}
