// File to print the instance of the structs

package internal

import "fmt"

func (person *Person) PrintName() {
	fmt.Println("The name of the person is:", person.FirstNmame)
}

func (animal *Animal) PrintName() {
	fmt.Println("The name of the animal is:", animal.Name)
}
