package main

import (
	"github.com/aws/aws-lambda-go/lambda"
	"github.com/sbstjn/appsync-router"
)

// PeopleEvent matches the payload of the people resolver
type PeopleEvent struct{}

// PersonEvent matches the payload of the person(id: Number) resolver
type PersonEvent struct {
	ID int `json:"id"`
}

func handlePeople(args PeopleEvent) (interface{}, error) {
	return people, nil
}

func handlePerson(args PersonEvent) (interface{}, error) {
	return people.ByID(args.ID)
}

var (
	r = router.New()
)

func init() {
	r.Add("people", handlePeople)
	r.Add("person", handlePerson)
}

func main() {
	lambda.Start(r.Handle)
}
