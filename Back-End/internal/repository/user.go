package repository

import (
	"context"
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
func (u *userRepository) Delete(ctx context.Context, id string) error {
	return u.db.WithContext(ctx).Where("id = ?", id).Update("deleted_at", gorm.Expr("NOW()")).Error
}

// HardDelete implements [model.UserRepository].
func (u *userRepository) HardDelete(ctx context.Context, id string) error {
	return u.db.WithContext(ctx).Where("id = ?", id).Delete(&model.User{}).Error
}

// FindByEmail implements [model.UserRepository].
func (u *userRepository) FindByEmail(ctx context.Context, email string) (*model.User, error) {
	var user model.User

	err := u.db.WithContext(ctx).
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
func (u *userRepository) FindByID(ctx context.Context, id string) (*model.User, error) {
	var user model.User

	err := u.db.WithContext(ctx).
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
func (u *userRepository) Store(ctx context.Context, user *model.User) error {
	return u.db.WithContext(ctx).Create(user).Error
}

// Update implements [model.UserRepository].
func (u *userRepository) Update(ctx context.Context, user *model.User) error {
	return u.db.WithContext(ctx).Save(user).Error
}
