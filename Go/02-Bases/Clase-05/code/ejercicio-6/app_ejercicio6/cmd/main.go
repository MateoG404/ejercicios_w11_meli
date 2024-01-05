package main

import (
	"ejercicio_6/app_ejercicio6/internal"
	"fmt"
)

func main() {

	// Input value for create the new employee, name and hours worked by month
	var name string
	var hoursWorkedByMonth float32
	var valueHour float32

	fmt.Println("Ingrese el nombre del empleado:")
	fmt.Scan(&name)
	fmt.Println("Ingrese las horas trabajadas por mes:")
	fmt.Scan(&hoursWorkedByMonth)
	fmt.Println("Ingrese el valor por hora:")
	fmt.Scan(&valueHour)

	// Creation of the employee
	employee := internal.CreateEmployee(name)

	// Calculation of the employee's salary
	salary, err := employee.CalculateSalary(hoursWorkedByMonth, valueHour)
	fmt.Println("El salario de tu empleado es:")
	if err != nil {
		fmt.Println("Error: ", err)
		return
	}
	fmt.Println(salary)
}
