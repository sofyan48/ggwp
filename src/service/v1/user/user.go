package user

import (
	"time"

	"github.com/jinzhu/copier"
	dbEntity "github.com/sofyan48/ggwp/src/entity/v1/db"
	httpEntity "github.com/sofyan48/ggwp/src/entity/v1/http"
	repository "github.com/sofyan48/ggwp/src/repository/v1/user"
	"github.com/sofyan48/ggwp/src/util/helper/crypto"
)

// V1UserService derivated method
type V1UserService struct {
	userRepository repository.UserRepositoryInterface
}

// V1UserServiceHandler handler
// return: UserService
func V1UserServiceHandler() *V1UserService {
	return &V1UserService{
		userRepository: repository.UserRepositoryHandler(),
	}
}

// V1UserServiceInterface interface
type V1UserServiceInterface interface {
	UpdateUserByID(id int, payload httpEntity.UserRequest) bool
	InsertUsers(payload httpEntity.UserRequest) (*httpEntity.UserDetailResponse, error)
}

// UpdateUserByID params
// @id: int
// @payload: entity UserRequest
// return boolean
func (service *V1UserService) UpdateUserByID(id int, payload httpEntity.UserRequest) bool {
	var nows = time.Now()
	user := &dbEntity.Users{}
	user.Name = payload.Name
	user.Email = payload.Email
	user.Password = payload.Password
	user.DateOfBirth = payload.DateOfBirth
	user.PhoneNumber = payload.PhoneNumber
	user.CurrentAddress = payload.CurrentAddress
	user.City = payload.City
	user.Province = payload.Province
	user.District = payload.District
	user.Lat = payload.Lat
	user.Lng = payload.Lng
	user.Job = payload.Job
	user.Image = payload.Image
	user.UpdatedAt = &nows
	err := service.userRepository.UpdateUserByID(id, user)
	if nil != err {
		return false
	}
	return true
}

// InsertUsers params
// @id: int
// @payload: entity UserRequest
// return *httpEntity.UserDetailResponse, error
func (service *V1UserService) InsertUsers(payload httpEntity.UserRequest) (*httpEntity.UserDetailResponse, error) {
	var nows = time.Now()
	user := &dbEntity.Users{}
	user.Name = payload.Name
	user.Email = payload.Email
	user.Password, _ = crypto.HashPassword(payload.Password)
	user.DateOfBirth = payload.DateOfBirth
	user.PhoneNumber = payload.PhoneNumber
	user.CurrentAddress = payload.CurrentAddress
	user.City = payload.City
	user.Province = payload.Province
	user.District = payload.District
	user.Lat = payload.Lat
	user.Lng = payload.Lng
	user.Job = payload.Job
	user.Image = payload.Image
	user.UpdatedAt = &nows
	err := service.userRepository.InsertUsers(user)
	if err != nil {
		return nil, err
	}
	results := &httpEntity.UserDetailResponse{}
	copier.Copy(results, user)
	return results, nil
}
