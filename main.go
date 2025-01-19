package main

import (
	"employee-managment/algorithms"
	"employee-managment/models"
	"employee-managment/patterns"
	"employee-managment/services"
	"fmt"
)

//TIP <p>To run your code, right-click the code and select <b>Run</b>.</p> <p>Alternatively, click
// the <icon src="AllIcons.Actions.Execute"/> icon in the gutter and select the <b>Run</b> menu item from here.</p>

func main() {
	db := patterns.GetDatabaseInstance()
	fmt.Println(db)

	// Factory: Creating Employees
	fullTimeEmployee := patterns.EmployeeFactory(patterns.FullTimeEmployee, "1", "John", "Engineering", 30, 5000.0)
	partTimeEmployee := patterns.EmployeeFactory(patterns.PartTimeEmployee, "2", "Alice", "Marketing", 25, 3000.0)

	employeeService := services.NewEmployeeService()
	employeeService.AddEmployee(fullTimeEmployee.(*models.Employee))
	employeeService.AddEmployee(partTimeEmployee.(*models.Employee))

	// Strategy: Calculating Salaries
	fullTimeSalary := &patterns.FullTimeSalary{}
	partTimeSalary := &patterns.PartTimeSalary{}

	fmt.Printf("Full-time salary: %.2f\n", fullTimeSalary.CalculateSalary(fullTimeEmployee.(*models.Employee)))
	fmt.Printf("Part-time salary: %.2f\n", partTimeSalary.CalculateSalary(partTimeEmployee.(*models.Employee)))

	// Search: Finding an employee by name
	employees := employeeService.GetAllEmployees()
	foundEmployee := algorithms.LinearSearch(employees, "Alice")
	if foundEmployee != nil {
		fmt.Printf("Found Employee: %v\n", foundEmployee)
	}

	// Sort: Sorting employees by salary
	sortedEmployees := algorithms.SortBySalary(employees)
	fmt.Println("Sorted Employees by Salary:")
	for _, emp := range sortedEmployees {
		fmt.Printf("Name: %s, Salary: %.2f\n", emp.Name, emp.Salary)
	}
}
