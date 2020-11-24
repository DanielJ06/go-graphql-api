package main

import (
	"encoding/json"
	"fmt"
	"log"

	"github.com/graphql-go/graphql"
)

type Book struct {
	ID	int
	Title	string
	Author	Author
	Reviews	[]Review
}

type Author struct {
	Name	string
	Books	[]int
}

type Review struct {
	Body	string
}

func populate() []Book {
	author := &Author{Name: "Sidney Sheldon", Books: []int{1}}
	book := Book{
		ID: 1,
		Title: "The Other Side of Midnight",
		Author: *author,
		Reviews: []Review{
			Review{Body: "My favorite book"},
		},
	}

	var books []Book
	books = append(books, book)

	return books
}

func main() {

	var reviewType = graphql.NewObject(
		graphql.ObjectConfig{
			Name: "Review",
			Fields: graphql.Fields{
				"body": &graphql.Field{
					Type: graphql.String,
				},
			},
		},
	)

	var authorType = graphql.NewObject(
		graphql.ObjectConfig{
			Name: "Author",
			Fields: graphql.Fields{
				"Name": &graphql.Field{
					Type: graphql.String,
				},
				"Books": &graphql.Field{
					Type: graphql.NewList(graphql.Int),
				},
			},
		},
	)

	var bookType = graphql.NewObject(
		graphql.ObjectConfig{
			Name: "Book",
			Fields: graphql.Fields{
				"ID": &graphql.Field{
					Type: graphql.Int,
				},
				"Title": &graphql.Field{
					Type: graphql.String,
				},
				"Author": &graphql.Field{
					Type: authorType,
				},
				"Reviews": &graphql.Field{
					Type: graphql.NewList(reviewType),
				},
			},
		},
	)

	fields := graphql.Fields{
		"hello": &graphql.Field{
			Type: graphql.String,
			Resolve: func(p graphql.ResolveParams) (interface{}, error) {
				return "Hello world", nil
			},
		},
	}

	rootQuery := graphql.ObjectConfig{Name: "RootQuery", Fields: fields}
	schemaConfig := graphql.SchemaConfig{Query: graphql.NewObject(rootQuery)}
	schema, err := graphql.NewSchema(schemaConfig)
	if err != nil {
		log.Fatalf("Failed to create new gql schema, err %v", err)
	}

	query := `
		{
			hello
		}	
	`

	params := graphql.Params{Schema: schema, RequestString: query}
	r := graphql.Do(params) 
	if len(r.Errors) > 0 {
		log.Fatalf("Failed to execute graphql operation, errors? %+v", r.Errors)
	} 
	rJSON, _ := json.Marshal(r)
	fmt.Printf("%s \n", rJSON)

}