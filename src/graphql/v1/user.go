package v1

import (
	gql "github.com/graphql-go/graphql"
	types "github.com/sofyan48/ggwp/src/entity/v1/graphql"
	"github.com/sofyan48/ggwp/src/service/v1/user"
)

var userService = user.UserServiceHandler()

// Query fields
var Query = gql.Fields{
	"User": &gql.Field{
		Type: types.User,
		Args: gql.FieldConfigArgument{
			"page": &gql.ArgumentConfig{
				Type: gql.Int,
			},
			"limit": &gql.ArgumentConfig{
				Type: gql.Int,
			},
		},
		Resolve: userService.GetAllUser,
	},
}
