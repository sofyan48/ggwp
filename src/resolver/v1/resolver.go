package v1

import (
	gql "github.com/graphql-go/graphql"
	types "github.com/sofyan48/ggwp/src/entity/v1/graphql"
	user "github.com/sofyan48/ggwp/src/service/v1/graphql"
)

var userService = user.GraphQLUserServiceHandler()

// RootResolver fields
var RootResolver = gql.Fields{
	"Users": &gql.Field{
		Type: types.Users,
		Args: gql.FieldConfigArgument{
			"page": &gql.ArgumentConfig{
				Type: gql.Int,
			},
			"limit": &gql.ArgumentConfig{
				Type: gql.Int,
			},
		},
		Resolve: userService.GraphQLGetAllUser,
	},

	"User": &gql.Field{
		Type: types.User,
		Args: gql.FieldConfigArgument{
			"id": &gql.ArgumentConfig{
				Type: gql.Int,
			},
		},
		Resolve: userService.GraphQLGetUserByID,
	},
}
