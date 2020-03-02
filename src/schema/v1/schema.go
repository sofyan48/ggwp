package v1

import (
	gql "github.com/graphql-go/graphql"
	v1 "github.com/sofyan48/ggwp/src/graphql/v1"
)

// Schema config
func Schema() (gql.Schema, error) {
	schema, error := gql.NewSchema(gql.SchemaConfig{
		Query: gql.NewObject(gql.ObjectConfig{
			Name:   "Query",
			Fields: v1.Query,
		}),
	})

	return schema, error
}
