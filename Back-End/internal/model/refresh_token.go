package model

import (
	"context"
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type RefreshToken struct {
	ID        string     `gorm:"column:id;primaryKey" json:"id"`
	UserID    string     `gorm:"column:user_id" json:"user_id"`
	TokenHash string     `gorm:"column:token_hash;unique" json:"token"`
	ExpiresAt time.Time  `gorm:"column:expires_at" json:"expires_at"`
	RevokedAt *time.Time `gorm:"column:revoked_at" json:"revoked_at"`
	CreatedAt time.Time  `gorm:"column:created_at" json:"created_at"`
}

func (rt *RefreshToken) TableName() string {
	return "refresh_tokens"
}

func (rt *RefreshToken) BeforeCreate(tx *gorm.DB) (err error) {
	rt.ID = uuid.New().String()
	return nil
}

type RefreshTokenRepository interface {
	Store(ctx context.Context, token *RefreshToken) error
	FindByTokenHash(ctx context.Context, tokenHash string) (*RefreshToken, error)
	FindByUserID(ctx context.Context, userID string) ([]*RefreshToken, error)
	FindActiveByUserID(ctx context.Context, userID string) ([]*RefreshToken, error)
	Revoke(ctx context.Context, id string) error
}
