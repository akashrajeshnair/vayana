package repository

import (
	"github.com/akashrajeshnair/vayana/pkg/errors"
	"github.com/akashrajeshnair/vayana/services/user/internal/models"
	"gorm.io/gorm"
)

type UserRepository struct {
	db *gorm.DB
}

func NewUserRepository(db *gorm.DB) *UserRepository {
	return &UserRepository{db}
}

func (r *UserRepository) CreateUser(email, password, name string) (*models.User, error) {
	user := &models.User{
		Email:    email,
		Password: password,
		Name:     name,
	}

	if err := r.db.Create(user).Error; err != nil {
		return nil, errors.NewDatabaseError("failed to create new user", err)
	}

	return user, nil
}

func (r *UserRepository) GetUserByEmail(email string) (*models.User, error) {
	user := &models.User{}

	if err := r.db.Where("email = ?", email).First(user).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			return nil, errors.NewNotFoundError("user not found")
		}
		return nil, errors.NewDatabaseError("failed to get user by email", err)
	}

	return user, nil
}

func (r *UserRepository) GetUserByID(id string) (*models.User, error) {
	user := &models.User{}

	if err := r.db.Where("id = ?", id).First(user).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			return nil, errors.NewNotFoundError("user not found")
		}
		return nil, errors.NewDatabaseError("failed to get user by id", err)
	}

	return user, nil
}

func (r *UserRepository) UpdateUser(user *models.User) error {
	if err := r.db.Save(user).Error; err != nil {
		return errors.NewDatabaseError("failed to update user", err)
	}
	return nil
}

func (r *UserRepository) DeleteUser(id string) error {
	if err := r.db.Where("id = ?", id).Delete(&models.User{}).Error; err != nil {
		return errors.NewDatabaseError("failed to delete user", err)
	}
	return nil
}

func (r *UserRepository) GetAllUsers() ([]models.User, error) {
	var users []models.User
	if err := r.db.Find(&users).Error; err != nil {
		return nil, errors.NewDatabaseError("failed to get all users", err)
	}
	return users, nil
}

func (r *UserRepository) EmailExists(email string) (bool, error) {
	var count int64
	if err := r.db.Model(&models.User{}).Where("email = ?", email).Count(&count).Error; err != nil {
		return false, errors.NewDatabaseError("failed to check email existence", err)
	}
	return count > 0, nil
}
