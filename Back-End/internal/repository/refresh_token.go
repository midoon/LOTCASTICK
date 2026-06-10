package repository

import (
	"context"
	"lotcastick-backend/internal/model"
	"time"

	"gorm.io/gorm"
)

type refreshTokenRepository struct {
	db *gorm.DB
}

// FindActiveByUserID implements [model.RefreshTokenRepository].
func (r *refreshTokenRepository) FindActiveByUserID(ctx context.Context, userID string) ([]*model.RefreshToken, error) {
	var tokens []*model.RefreshToken
	err := r.db.WithContext(ctx).Where("user_id = ? AND revoked_at IS NULL AND expires_at > ?", userID, time.Now()).Find(&tokens).Error
	if err != nil {
		if err == gorm.ErrRecordNotFound {
			return nil, nil
		}
		return nil, err
	}
	return tokens, err
}

// FindByUserID implements [model.RefreshTokenRepository].
func (r *refreshTokenRepository) FindByUserID(ctx context.Context, userID string) ([]*model.RefreshToken, error) {
	var tokens []*model.RefreshToken
	err := r.db.WithContext(ctx).Where("user_id = ?", userID).Find(&tokens).Error
	if err != nil {
		if err == gorm.ErrRecordNotFound {
			return nil, nil
		}
		return nil, err
	}
	return tokens, err
}

// FindByTokenHash implements [model.RefreshTokenRepository].
func (r *refreshTokenRepository) FindByTokenHash(ctx context.Context, tokenHash string) (*model.RefreshToken, error) {
	var token model.RefreshToken
	err := r.db.WithContext(ctx).Where("token_hash = ?", tokenHash).First(&token).Error
	if err != nil {
		if err == gorm.ErrRecordNotFound {
			return nil, nil
		}
		return nil, err
	}
	return &token, nil
}

// Revoke implements [model.RefreshTokenRepository].
func (r *refreshTokenRepository) Revoke(ctx context.Context, id string) error {
	return r.db.WithContext(ctx).Model(&model.RefreshToken{}).Where("id = ?", id).Update("revoked_at", time.Now()).Error
}

// Store implements [model.RefreshTokenRepository].
func (r *refreshTokenRepository) Store(ctx context.Context, token *model.RefreshToken) error {
	return r.db.WithContext(ctx).Create(token).Error
}

func NewRefreshTokenRepository(db *gorm.DB) model.RefreshTokenRepository {
	return &refreshTokenRepository{db: db}
}
