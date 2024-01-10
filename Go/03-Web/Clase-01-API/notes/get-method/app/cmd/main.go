package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"

	"github.com/go-chi/chi/v5"
)

// Person representa una persona
type Person struct {
	Name string `json:"name"`
	Age  int    `json:"age"`
}

// PeopleRepository es un repositorio de personas
type PeopleRepository struct {
	People map[int]Person
}

// GetPeople devuelve todas las personas en el repositorio
func (pr *PeopleRepository) GetPeople() ([]Person, error) {
	people := make([]Person, 0, len(pr.People))
	for _, v := range pr.People {
		people = append(people, v)
	}
	return people, nil
}

func main() {
	// Inicializar el router
	r := chi.NewRouter()

	// Inicializar el repositorio de personas con valores hardcoded
	peopleRepo := &PeopleRepository{
		People: map[int]Person{
			1: {Name: "John", Age: 30},
			2: {Name: "Jane", Age: 25},
			3: {Name: "Bob", Age: 25},
			4: {Name: "Alice", Age: 28},
		},
	}

	r.Get("/people", func(w http.ResponseWriter, r *http.Request) {
		// Obtener query parameters
		nameFilter := r.URL.Query().Get("name")
		ageFilterStr := r.URL.Query().Get("age")

		// Filtrar por nombre si se proporciona
		var filteredPeople []Person
		if nameFilter != "" {
			for _, person := range peopleRepo.People {
				if person.Name == nameFilter {
					filteredPeople = append(filteredPeople, person)
				}
			}
		} else {
			// Si no se proporciona un filtro de nombre, mostrar todas las personas
			people, err := peopleRepo.GetPeople()
			if err != nil {
				http.Error(w, err.Error(), http.StatusInternalServerError)
				return
			}
			filteredPeople = people
		}

		// Filtrar por edad si se proporciona
		if ageFilterStr != "" {
			ageFilter, err := strconv.Atoi(ageFilterStr)
			if err != nil {
				http.Error(w, "Invalid age parameter", http.StatusBadRequest)
				return
			}
			var ageFilteredPeople []Person
			for _, person := range filteredPeople {
				if person.Age == ageFilter {
					ageFilteredPeople = append(ageFilteredPeople, person)
				}
			}
			filteredPeople = ageFilteredPeople
		}

		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(filteredPeople)
	})

	fmt.Println("Server listening on :8080")
	http.ListenAndServe(":8080", r)
}
