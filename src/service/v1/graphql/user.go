package user

import (
	"errors"

	"github.com/graphql-go/graphql"
	dbEntity "github.com/sofyan48/ggwp/src/entity/v1/db"
	repository "github.com/sofyan48/ggwp/src/repository/v1/user"
)

// GraphQLUserService derivated method
type GraphQLUserService struct {
	userRepository repository.UserRepositoryInterface
}

// GraphQLUserServiceHandler handler
// return: UserService
func GraphQLUserServiceHandler() *GraphQLUserService {
	return &GraphQLUserService{
		userRepository: repository.UserRepositoryHandler(),
	}
}

type GraphQLUserServiceInterface interface {
	GraphQLGetUserByID(gq graphql.ResolveParams) (interface{}, error)
	GraphQLGetAllUser(gq graphql.ResolveParams) (interface{}, error)
}

// GraphQLGetUserByID params
// @id: int
// @waitGroup: *sync.WaitGroup
// return interface{}, error
func (service *GraphQLUserService) GraphQLGetUserByID(gq graphql.ResolveParams) (interface{}, error) {
	id, hasID := gq.Args["id"].(int)
	if !hasID {
		return "", errors.New("ID Not Valid")
	}
	userData := &dbEntity.Users{}
	result, err := service.userRepository.GetUserByID(id, userData)

	if err != nil {
		return "", errors.New("Function Errors")
	}
	return result, nil
}

// GraphQLGetAllUser params
// @id: int
// @count: int
// return interface{}, error
func (service *GraphQLUserService) GraphQLGetAllUser(gq graphql.ResolveParams) (interface{}, error) {
	page, hasPage := gq.Args["page"].(int)
	limit, hasLimit := gq.Args["limit"].(int)
	if !hasLimit && !hasPage {
		return "", errors.New("Limit And Page Not Resolve")
	}
	users, _ := service.userRepository.GetUsersList(limit, page)
	result := map[string]interface{}{
		"results": users,
		"page":    page,
	}
	return result, nil
}
