package graphql

import (
	gql "github.com/graphql-go/graphql"
)

// User fields
var User = gql.NewObject(gql.ObjectConfig{
	Name: "User",
	Fields: gql.Fields{
		"ID": &gql.Field{
			Type: gql.Int,
		},
		"name": &gql.Field{
			Type: gql.String,
		},
		"email": &gql.Field{
			Type: gql.String,
		},
		"password": &gql.Field{
			Type: gql.String,
		},
		"dateofbirth": &gql.Field{
			Type: gql.String,
		},
		"phoneNumber": &gql.Field{
			Type: gql.String,
		},
		"currentAddress": &gql.Field{
			Type: gql.String,
		},
		"city": &gql.Field{
			Type: gql.String,
		},
		"province": &gql.Field{
			Type: gql.String,
		},
		"district": &gql.Field{
			Type: gql.String,
		},
		"lat": &gql.Field{
			Type: gql.String,
		},
		"lng": &gql.Field{
			Type: gql.String,
		},
		"job": &gql.Field{
			Type: gql.String,
		},
		"image": &gql.Field{
			Type: gql.String,
		},
	},
})
