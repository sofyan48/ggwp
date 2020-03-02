package user

import (
	"errors"

	"github.com/graphql-go/graphql"
	"github.com/jinzhu/copier"
	dbEntity "github.com/sofyan48/ggwp/src/entity/v1/db"
	httpEntity "github.com/sofyan48/ggwp/src/entity/v1/http"
	repository "github.com/sofyan48/ggwp/src/repository/v1/user"
)

// UserService derivated method
type V1UserService struct {
	userRepository repository.UserRepositoryInterface
}

// UserServiceHandler handler
// return: UserService
func UserServiceHandler() *V1UserService {
	return &V1UserService{
		userRepository: repository.UserRepositoryHandler(),
	}
}

// GetUserByID params
// @id: int
// @waitGroup: *sync.WaitGroup
// return entity UserDetailResponse
func (service *V1UserService) GetUserByID(gq graphql.ResolveParams) (interface{}, error) {
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

// GetAllUser params
// @id: int
// @count: int
// return entity UserResponse
func (service *V1UserService) GetAllUser(gq graphql.ResolveParams) (interface{}, error) {
	data := []httpEntity.UserResponse{}
	page, hasPage := gq.Args["page"].(int)
	limit, hasLimit := gq.Args["limit"].(int)
	if !hasLimit && !hasPage {
		return "", errors.New("Limit And Page Not Resolve")
	}
	users, _ := service.userRepository.GetUsersList(limit, page)
	copier.Copy(&data, &users)
	result := map[string]interface{}{
		"results": data,
		"page":    page,
	}
	return result, nil
}
