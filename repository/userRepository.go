package repository

import (
	"errors"
	"kiss_web/models"

	"gorm.io/gorm"
)

// IUserRepository ...
type IUserRepository interface {
	CreateUser(user *models.User) error
	UpdateUser(user *models.User) error
	GetUserByID(id uint64) (*models.User, error)
	GetAllUsers() (*[]models.User, error)

	RemoveUser(id uint64) error
}

// UserRepository ...
type UserRepository struct {
	db *gorm.DB
}

// NewUserRepository ....
func NewUserRepository(connection *gorm.DB) IUserRepository {
	return &UserRepository{
		db: connection,
	}
}

//CreateUser user to db.Table("users")
func (repo *UserRepository) CreateUser(user *models.User) error {

	var count int64
	repo.db.Table("users").Where("login = ?", user.Login).Count(&count)

	if count > 0 {
		return errors.New("A user with the same email address already exists")
	}

	result := repo.db.Create(user)

	if result.Error != nil {
		return result.Error
	}

	return nil
}

// RemoveUser delete user by ID...
func (repo *UserRepository) RemoveUser(id uint64) error {

	result := repo.db.Delete(&models.User{ID: id})

	if result.Error != nil {
		return result.Error
	}

	if result.RowsAffected > 0 {
		return nil
	}

	return errors.New("User not found")
}

// UpdateUser ...
func (repo *UserRepository) UpdateUser(user *models.User) error {

	var newUser models.User
	result := repo.db.First(&newUser, user.ID)
	if result.Error != nil {
		return result.Error
	}

	newUser.Login = user.Login
	newUser.Password = user.Password

	result = repo.db.Save(newUser)
	return result.Error
}

// GetUserByID ...
// return first user by ID
func (repo *UserRepository) GetUserByID(id uint64) (*models.User, error) {
	user := &models.User{}
	result := repo.db.First(user, id)
	return user, result.Error
}

// GetAllUsers return all users from db table 'users' // no public!
func (repo *UserRepository) GetAllUsers() (*[]models.User, error) {
	var users []models.User
	result := repo.db.Table("users").Scan(&users)
	return &users, result.Error
}
