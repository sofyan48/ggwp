package v1

import (
	gql "github.com/graphql-go/graphql"
	v1 "github.com/sofyan48/ggwp/src/resolver/v1"
)

// Schema config
func Schema() (gql.Schema, error) {
	schema, error := gql.NewSchema(gql.SchemaConfig{
		Query: gql.NewObject(gql.ObjectConfig{
			Name:   "Query",
			Fields: v1.RootResolver,
		}),
	})

	return schema, error
}
