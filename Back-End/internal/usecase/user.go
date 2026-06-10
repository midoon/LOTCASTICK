package usecase

import (
	"context"
	"lotcastick-backend/internal/dto"
	"lotcastick-backend/internal/model"
	"lotcastick-backend/internal/util"
	"net/http"
	"time"

	"github.com/go-playground/validator/v10"
	"github.com/spf13/viper"
)

type userUsecase struct {
	userRepo    model.UserRepository
	tokenRepo   model.RefreshTokenRepository
	validate    *validator.Validate
	viperConfig *viper.Viper
}

func NewUserUsecase(userRepo model.UserRepository, tokenRepo model.RefreshTokenRepository, validate *validator.Validate, viperConfig *viper.Viper) model.UserUsecase {
	return &userUsecase{userRepo: userRepo, tokenRepo: tokenRepo, validate: validate, viperConfig: viperConfig}
}

// Register implements [model.UserUsecase].
func (u *userUsecase) Register(ctx context.Context, req dto.RegisterRequest) error {
	if err := u.validate.Struct(req); err != nil {
		return util.NewCustomError(http.StatusBadRequest, "validation error", err)
	}

	isExisting, err := u.userRepo.FindByEmail(ctx, req.Email)
	if err != nil {
		return util.NewCustomError(http.StatusInternalServerError, "internal server error", err)
	}

	if isExisting != nil {
		return util.NewCustomError(http.StatusConflict, "email already exists", nil)
	}

	hashedPassword, err := util.CreateHash(req.Password)
	if err != nil {
		return util.NewCustomError(http.StatusInternalServerError, "internal server error", err)
	}
	user := &model.User{
		Email:           req.Email,
		PasswordHash:    hashedPassword,
		DisplayName:     req.DisplayName,
		Timezone:        req.Timezone,
		DefaultCurrency: req.DefaultCurrency,
	}

	if err := u.userRepo.Store(ctx, user); err != nil {
		return util.NewCustomError(http.StatusInternalServerError, "internal server error", err)
	}
	return nil
}

// Login implements [model.UserUsecase].
func (u *userUsecase) Login(ctx context.Context, req dto.LoginRequest) (*dto.TokenData, error) {
	if err := u.validate.Struct(req); err != nil {
		return nil, util.NewCustomError(http.StatusBadRequest, "validation error", err)
	}

	//is user exists by email
	user, err := u.userRepo.FindByEmail(ctx, req.Email)
	if err != nil {
		return nil, util.NewCustomError(http.StatusInternalServerError, "internal server error", err)
	}
	if user == nil {
		return nil, util.NewCustomError(http.StatusUnauthorized, "invalid credentials", nil)
	}
	if !util.CompareHash(user.PasswordHash, req.Password) {
		return nil, util.NewCustomError(http.StatusUnauthorized, "invalid credentials", nil)
	}

	// is JWT exist in DB
	tokenExisting, err := u.tokenRepo.FindActiveByUserID(ctx, user.ID)
	if err != nil {
		return nil, util.NewCustomError(http.StatusInternalServerError, "internal server error", err)
	}

	if len(tokenExisting) > 0 {
		for _, t := range tokenExisting {
			if err := u.tokenRepo.Revoke(ctx, t.ID); err != nil {
				return nil, util.NewCustomError(http.StatusInternalServerError, "internal server error", err)
			}
		}
	}

	aToken, err := util.GenerateAccessToken(user.ID, time.Duration(u.viperConfig.GetInt64("jwt.expiration"))*time.Second, u.viperConfig.GetString("jwt.secret"))
	if err != nil {
		return nil, util.NewCustomError(http.StatusInternalServerError, "internal server error", err)
	}

	rToken, err := util.GenerateRefreshToken(user.ID, time.Duration(u.viperConfig.GetInt64("jwt.refresh_expiration"))*time.Second, u.viperConfig.GetString("jwt.secret"))
	if err != nil {
		return nil, util.NewCustomError(http.StatusInternalServerError, "internal server error", err)
	}

	tokenHash := util.CreateTokenHash(rToken)

	refreshToken := &model.RefreshToken{
		UserID:    user.ID,
		TokenHash: tokenHash,
		ExpiresAt: time.Now().Add(time.Duration(u.viperConfig.GetInt64("jwt.refresh_expiration")) * time.Second),
	}

	err = u.tokenRepo.Store(ctx, refreshToken)
	if err != nil {
		return nil, util.NewCustomError(http.StatusInternalServerError, "internal server error", err)
	}

	return &dto.TokenData{
		UserID:       user.ID,
		AccessToken:  aToken,
		RefreshToken: rToken,
		ExpiresIn:    u.viperConfig.GetInt64("jwt.expiration"),
	}, nil

}

func (u *userUsecase) Logout(ctx context.Context, userID string) error {
	// Revoke all active tokens for the user
	tokenExisting, err := u.tokenRepo.FindActiveByUserID(ctx, userID)
	if err != nil {
		return util.NewCustomError(http.StatusInternalServerError, "internal server error", err)
	}
	for _, t := range tokenExisting {
		if err := u.tokenRepo.Revoke(ctx, t.ID); err != nil {
			return util.NewCustomError(http.StatusInternalServerError, "internal server error", err)
		}
	}
	return nil
}

// DeleteAccount implements [model.UserUsecase].
func (u *userUsecase) DeleteAccount(ctx context.Context, userID string) error {
	panic("unimplemented")
}

// GetProfile implements [model.UserUsecase].
func (u *userUsecase) GetProfile(ctx context.Context, userID string) (*model.User, error) {
	panic("unimplemented")
}

// UpdateProfile implements [model.UserUsecase].
func (u *userUsecase) UpdateProfile(ctx context.Context, userID string, displayName string, timezone string, defaultCurrency string) (*model.User, error) {
	panic("unimplemented")
}
