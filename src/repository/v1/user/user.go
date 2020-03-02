package user

import (
	"github.com/jinzhu/gorm"

	dbEntity "github.com/sofyan48/ggwp/src/entity/v1/db"
	connection "github.com/sofyan48/ggwp/src/util/helper/mysqlconnection"
)

// UserRepository types
type UserRepository struct {
	DB gorm.DB
}

// UserRepositoryHandler Users handler repo
// return: UserRepository
func UserRepositoryHandler() *UserRepository {
	return &UserRepository{
		DB: *connection.GetConnection(),
	}
}

// UserRepositoryInterface interface
type UserRepositoryInterface interface {
	GetUserByID(id int, userData *dbEntity.Users) (*dbEntity.Users, error)
	UpdateUserByID(id int, userData *dbEntity.Users) error
	GetUsersList(limit int, offset int) ([]dbEntity.Users, error)
	InsertUsers(userData *dbEntity.Users) error
	CheckEmailUsers(email string, usersData *dbEntity.Users) bool
}

// GetUserByID params
// @id: int
// @userData: entity Users
// wg *sync.WaitGroup
// return error
func (repository *UserRepository) GetUserByID(id int, userData *dbEntity.Users) (*dbEntity.Users, error) {
	query := repository.DB.Table("users")
	query = query.Where("id=?", id)
	query = query.First(&userData)
	return userData, query.Error
}

// UpdateUserByID params
// @id: int
// @userData: entity Users
// return error
func (repository *UserRepository) UpdateUserByID(id int, userData *dbEntity.Users) error {
	query := repository.DB.Table("users")
	query = query.Where("id=?", id)
	query = query.Updates(userData)
	query.Scan(&userData)
	return query.Error
}

// GetUsersList params
// @id: int
// @userData: entity Users
// return entity,error
func (repository *UserRepository) GetUsersList(limit int, offset int) ([]dbEntity.Users, error) {
	users := []dbEntity.Users{}
	query := repository.DB.Table("users")
	query = query.Limit(10).Offset(0)
	query = query.Find(&users)
	return users, query.Error
}

// InsertUsers params
// @userData: entity Users
// return error
func (repository *UserRepository) InsertUsers(usersData *dbEntity.Users) error {
	query := repository.DB.Table("users")
	query = query.Create(usersData)
	query.Scan(&usersData)
	return query.Error
}

// CheckEmailUsers params
// @email : string
// @userData: entity Users
// return error
func (repository *UserRepository) CheckEmailUsers(email string, usersData *dbEntity.Users) bool {
	query := repository.DB.Table("users")
	if err := query.Where("email=?", email).First(&usersData).Error; err != nil {
		return false
	}
	return true
}
