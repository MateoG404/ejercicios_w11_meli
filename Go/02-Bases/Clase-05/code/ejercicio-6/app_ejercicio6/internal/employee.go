// This file contains the employee struct and its methods to manage the employee's salary
package internal

import (
	"errors"
)

// Employee struct to manage the employee's salary
type Employee struct {
	name               string
	hoursWorkedByMonth float32
	valueHour          float32
	salary             float64
}

// Methods

// CreateEmployee() to create a new employee
func CreateEmployee(name string) *Employee {
	return &Employee{name: name}
}

// CalculateSalary() to calculate the employee's salary
func (e *Employee) CalculateSalary(hoursWorkedByMonth float32, valueHour float32) (float64, error) {

	e.salary = float64(hoursWorkedByMonth * valueHour)
	if hoursWorkedByMonth < 80 {
		return 0, errors.New("The worker cannot have worked less that 80 hours per month")
	} else if e.salary > 150000 {
		return e.salary - (e.salary * 0.1), nil
	} else if hoursWorkedByMonth >= 80 {
		return e.salary, nil
	}
	return 0, errors.New("Unknown error")
}
