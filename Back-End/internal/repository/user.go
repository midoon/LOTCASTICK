package repository

import (
	"errors"
	"lotcastick-backend/internal/model"

	"gorm.io/gorm"
)

type userRepository struct {
	db *gorm.DB
}

func NewUserRepository(db *gorm.DB) model.UserRepository {
	return &userRepository{db: db}
}

// Delete implements [model.UserRepository].
func (u *userRepository) Delete(id string) error {
	return u.db.Where("id = ?", id).Update("deleted_at", gorm.Expr("NOW()")).Error
}

// HardDelete implements [model.UserRepository].
func (u *userRepository) HardDelete(id string) error {
	return u.db.Where("id = ?", id).Delete(&model.User{}).Error
}

// FindByEmail implements [model.UserRepository].
func (u *userRepository) FindByEmail(email string) (*model.User, error) {
	var user model.User

	err := u.db.
		Where("email = ?", email).Where("deleted_at IS NULL").
		First(&user).
		Error

	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, nil
		}
		return nil, err
	}

	return &user, nil
}

// FindByID implements [model.UserRepository].
func (u *userRepository) FindByID(id string) (*model.User, error) {
	var user model.User

	err := u.db.
		Where("id = ?", id).Where("deleted_at IS NULL").
		First(&user).
		Error

	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, nil
		}
		return nil, err
	}

	return &user, nil
}

// Store implements [model.UserRepository].
func (u *userRepository) Store(user *model.User) error {
	return u.db.Create(user).Error
}

// Update implements [model.UserRepository].
func (u *userRepository) Update(user *model.User) error {
	return u.db.Save(user).Error
}
