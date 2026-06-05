package model

import "time"

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
	Store(user *User) error
	FindByEmail(email string) (*User, error)
	FindByID(id string) (*User, error)
	Update(user *User) error
	Delete(id string) error
	HardDelete(id string) error
}

type UserUsecase interface {
	Register(email, password, displayName, timezone, defaultCurrency string) (*User, error)
	Login(email, password string) (*User, error)
	GetProfile(userID string) (*User, error)
	UpdateProfile(userID, displayName, timezone, defaultCurrency string) (*User, error)
	DeleteAccount(userID string) error
}
