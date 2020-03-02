package user

import (
	"errors"
	"sync"

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

// UserServiceInterface interface
type UserServiceInterface interface {
	GetUserByID(id int, waitGroup *sync.WaitGroup) *httpEntity.UserDetailResponse
	GetAllUser(page int, count int) []httpEntity.UserResponse
}

// GetUserByID params
// @id: int
// @waitGroup: *sync.WaitGroup
// return entity UserDetailResponse
func (service *V1UserService) GetUserByID(id int, waitGroup *sync.WaitGroup) *httpEntity.UserDetailResponse {
	userData := &dbEntity.Users{}
	result := &httpEntity.UserDetailResponse{}
	waitGroup.Add(1)
	go service.userRepository.GetUserByID(id, userData, waitGroup)
	waitGroup.Wait()
	copier.Copy(&result, &userData)
	return result
}

// GetAllUser params
// @id: int
// @count: int
// return entity UserResponse
func (service *V1UserService) GetAllUser(gq graphql.ResolveParams) (interface{}, error) {
	page, hasPage := gq.Args["page"].(int)
	limit, hasLimit := gq.Args["limit"].(int)
	if !hasLimit && !hasPage {
		return "", errors.New("Limit And Page Not Resolve")
	}
	users, _ := service.userRepository.GetUsersList(limit, page)
	result := []httpEntity.UserResponse{}
	copier.Copy(&result, &users)
	return result, nil
}
