package main

import (
	"fmt"
)

// Person is a person
type Person struct {
	ID   int    `json:"id"`
	Name string `json:"name"`
	Age  int    `json:"age"`
}

// People is a list of Person
type People []*Person

// ByID looks for a Person in People with the given ID
func (p People) ByID(id int) (*Person, error) {
	for _, person := range p {
		if person.ID == id {
			return person, nil
		}
	}

	return nil, fmt.Errorf("Cannot find person with ID: %d", id)
}

var (
	people = People{
		&Person{1, "Frank Ocean", 47},
		&Person{2, "Paul Gascoigne", 51},
		&Person{3, "Uwe Seeler", 81},
	}
)
