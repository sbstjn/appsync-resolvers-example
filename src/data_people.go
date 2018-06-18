package main

import "fmt"

type people []*person

func (p people) byID(id int) (*person, error) {
	for _, person := range p {
		if person.ID == id {
			return person, nil
		}
	}

	return nil, fmt.Errorf("Cannot find person with ID: %d", id)
}
