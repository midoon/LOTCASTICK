package model

import (
	"context"
	"lotcastick-backend/internal/dto"
	"time"
)

type User struct {
	ID              string     `gorm:"column:id;primaryKey" json:"id"`
	Email           string     `gorm:"column:email;unique" json:"email"`
	PasswordHash    string     `gorm:"column:password_hash" json:"-"`
	DisplayName     string     `gorm:"column:display_name" json:"display_name"`
	Timezone        string     `gorm:"column:timezone" json:"timezone"`
	DefaultCurrency string     `gorm:"column:default_currency" json:"default_currency"`
	LoginAttempts   int        `gorm:"column:login_attempts" json:"-"`
	LockedUntil     *time.Time `gorm:"column:locked_until" json:"-"`
	CreatedAt       *time.Time `gorm:"column:created_at" json:"created_at"`
	UpdatedAt       *time.Time `gorm:"column:updated_at" json:"updated_at"`
	DeletedAt       *time.Time `gorm:"column:deleted_at" json:"-"`
}

func (u *User) TableName() string {
	return "users"
}

type UserRepository interface {
	Store(ctx context.Context, user *User) error
	FindByEmail(ctx context.Context, email string) (*User, error)
	FindByID(ctx context.Context, id string) (*User, error)
	Update(ctx context.Context, user *User) error
	Delete(ctx context.Context, id string) error
	HardDelete(ctx context.Context, id string) error
}

type UserUsecase interface {
	Register(ctx context.Context, req dto.RegisterRequest) (*User, error)
	Login(ctx context.Context, email, password string) (*User, error)
	GetProfile(ctx context.Context, userID string) (*User, error)
	UpdateProfile(ctx context.Context, userID, displayName, timezone, defaultCurrency string) (*User, error)
	DeleteAccount(ctx context.Context, userID string) error
}
