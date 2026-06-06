package usecase

import (
	"context"
	"lotcastick-backend/internal/dto"
	"lotcastick-backend/internal/model"
	"lotcastick-backend/internal/util"
	"net/http"

	"github.com/go-playground/validator/v10"
)

type userUsecase struct {
	userRepo model.UserRepository
	validate *validator.Validate
}

func NewUserUsecase(userRepo model.UserRepository, validate *validator.Validate) model.UserUsecase {
	return &userUsecase{userRepo: userRepo, validate: validate}
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

// DeleteAccount implements [model.UserUsecase].
func (u *userUsecase) DeleteAccount(ctx context.Context, userID string) error {
	panic("unimplemented")
}

// GetProfile implements [model.UserUsecase].
func (u *userUsecase) GetProfile(ctx context.Context, userID string) (*model.User, error) {
	panic("unimplemented")
}

// Login implements [model.UserUsecase].
func (u *userUsecase) Login(ctx context.Context, email string, password string) (*model.User, error) {
	panic("unimplemented")
}

// UpdateProfile implements [model.UserUsecase].
func (u *userUsecase) UpdateProfile(ctx context.Context, userID string, displayName string, timezone string, defaultCurrency string) (*model.User, error) {
	panic("unimplemented")
}
