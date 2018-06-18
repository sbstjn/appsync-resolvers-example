package main

import (
	"github.com/aws/aws-lambda-go/lambda"
	resolvers "github.com/sbstjn/appsync-resolvers"
)

type personEvent struct {
	ID int `json:"id"`
}

func handlePeople() (interface{}, error) {
	return dataPeople, nil
}

func handlePerson(p personEvent) (interface{}, error) {
	return dataPeople.byID(p.ID)
}

func handleFriends(p person) (interface{}, error) {
	return p.getFriends()
}

var (
	r = resolvers.New()
)

func init() {
	r.Add("query.people", handlePeople)
	r.Add("query.person", handlePerson)
	r.Add("field.person.friends", handleFriends)
}

func main() {
	lambda.Start(r.Handle)
}
