package repository

import (
	"github.com/gesangwidigdo/auctify-be/interfaces"
	"github.com/gesangwidigdo/auctify-be/model"
	"gorm.io/gorm"
)

type userRepository struct {
	db *gorm.DB
}

func NewUserRepository(db *gorm.DB) interfaces.UserRepository {
	return &userRepository{
		db,
	}
}

// Login implements interfaces.UserRepository.
func (u *userRepository) GetByUsername(username string) (model.User, error) {
	var user model.User
	if err := u.db.Raw("SELECT id, username, password, role FROM users WHERE username = ? AND deleted_at IS NULL", username).Take(&user).Error; err != nil {
		return model.User{}, err
	}
	return user, nil
}

// Register implements interfaces.UserRepository.
func (u *userRepository) Register(request model.User) (model.User, error) {
	if err := u.db.Create(&request).Error; err != nil {
		return model.User{}, err
	}
	return request, nil
}

// Delete implements interfaces.UserRepository.
func (u *userRepository) Delete(id uint) error {
	if err := u.db.Exec("UPDATE users SET deleted_at = NOW() WHERE id = ? AND deleted_at IS NULL", id).Error; err != nil {
		return err
	}
	return nil
}

// Detail implements interfaces.UserRepository.
func (u *userRepository) Detail(id uint) (model.User, error) {
	var user model.User
	if err := u.db.Raw("SELECT id, name, email, username, address FROM users WHERE id = ? AND deleted_at IS NULL", id).Take(&user).Error; err != nil {
		return model.User{}, err
	}
	return user, nil
}

// List implements interfaces.UserRepository.
func (u *userRepository) List() ([]model.User, error) {
	var users []model.User
	if err := u.db.Raw("SELECT id, name, username, role FROM users WHERE deleted_at IS NULL").Scan(&users).Error; err != nil {
		return nil, err
	}
	return users, nil
}

// Update implements interfaces.UserRepository.
func (u *userRepository) Update(id uint, request model.User) (model.User, error) {
	if err := u.db.Model(&model.User{}).Where("id = ? AND deleted_at IS NULL", id).Updates(request).Error; err != nil {
		return model.User{}, err
	}
	return request, nil
}
