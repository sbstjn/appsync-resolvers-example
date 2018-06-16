package main

import (
	"github.com/aws/aws-lambda-go/lambda"
	resolvers "github.com/sbstjn/appsync-resolvers"
)

// PersonEvent matches the payload of the person(id: Number) resolver
type PersonEvent struct {
	ID int `json:"id"`
}

func handlePeople() (interface{}, error) {
	return people, nil
}

func handlePerson(args PersonEvent) (interface{}, error) {
	return people.ByID(args.ID)
}

var (
	r = resolvers.New()
)

func init() {
	r.Add("people", handlePeople)
	r.Add("person", handlePerson)
}

func main() {
	lambda.Start(r.Handle)
}
