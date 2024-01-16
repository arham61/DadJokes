package schemas

import (
	"errors"
	"github.com/graphql-go/graphql"
	"gqljokes/assets"
)

var jokeType = graphql.NewObject(
	graphql.ObjectConfig{
		Name: "Joke",
		Fields: graphql.Fields{
			"id":    &graphql.Field{Type: graphql.String},
			"title": &graphql.Field{Type: graphql.String},
			"text":  &graphql.Field{Type: graphql.String},
		},
	},
)

var rootQuery = graphql.NewObject(
	graphql.ObjectConfig{
		Name: "RootQuery",
		Fields: graphql.Fields{
			"joke": &graphql.Field{
				Type: jokeType,
				Args: graphql.FieldConfigArgument{
					"id": &graphql.ArgumentConfig{Type: graphql.String},
				},
				Resolve: getJokeByID,
			},
			"jokes": &graphql.Field{
				Type:    graphql.NewList(jokeType),
				Resolve: getAllJokes,
			},
		},
	},
)

var Schema, _ = graphql.NewSchema(
	graphql.SchemaConfig{
		Query: rootQuery,
	},
)

func getJokeByID(p graphql.ResolveParams) (interface{}, error) {
	id, ok := p.Args["id"].(string)
	if !ok {
		return nil, errors.New("ID argument is required")
	}

	for _, j := range assets.Jokes {
		if j.ID == id {
			return j, nil
		}
	}

	return nil, errors.New("Joke not found")
}

func getAllJokes(p graphql.ResolveParams) (interface{}, error) {
	return assets.Jokes, nil
}
