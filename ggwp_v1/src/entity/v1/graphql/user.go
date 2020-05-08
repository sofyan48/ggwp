package graphql

import (
	gql "github.com/graphql-go/graphql"
)

// ResultUsers fields
var ResultUsers = gql.NewObject(gql.ObjectConfig{
	Name: "results",
	Fields: gql.Fields{
		"id": &gql.Field{
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
	},
})

// Users type
var Users = gql.NewObject(gql.ObjectConfig{
	Name: "Users",
	Fields: gql.Fields{
		"results": &gql.Field{
			Type: gql.NewList(ResultUsers),
		},
		"page": &gql.Field{
			Type: gql.Int,
		},
	},
})

// User type
var User = gql.NewObject(gql.ObjectConfig{
	Name: "User",
	Fields: gql.Fields{
		"id": &gql.Field{
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
	},
})
