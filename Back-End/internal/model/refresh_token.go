package model

import "time"

type RefreshToken struct {
	ID        string     `gorm:"column:id;primaryKey" json:"id"`
	UserID    string     `gorm:"column:user_id" json:"user_id"`
	TokenHash string     `gorm:"column:token_hash;unique" json:"token"`
	ExpiresAt *time.Time `gorm:"column:expires_at" json:"expires_at"`
	RevokedAt *time.Time `gorm:"column:revoked_at" json:"revoked_at"`
	CreatedAt *time.Time `gorm:"column:created_at" json:"created_at"`
}

func (rt *RefreshToken) TableName() string {
	return "refresh_tokens"
}
