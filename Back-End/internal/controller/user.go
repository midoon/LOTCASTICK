package controller

import (
	"encoding/json"
	"lotcastick-backend/internal/dto"
	"lotcastick-backend/internal/model"
	"lotcastick-backend/internal/util"
	"net/http"
)

type UserController struct {
	userUsecase model.UserUsecase
}

func NewUserController(userUsecase model.UserUsecase) *UserController {
	return &UserController{userUsecase: userUsecase}
}

func (c *UserController) Register(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()

	request := dto.RegisterRequest{}
	if err := json.NewDecoder(r.Body).Decode(&request); err != nil {
		util.WriteJSON(w, http.StatusBadRequest, dto.ErrorResponse{
			Status:  false,
			Message: "Invalid request body",
			Error:   err.Error(),
		})
	}

	if err := c.userUsecase.Register(ctx, request); err != nil {
		if customErr, ok := err.(*util.CustomError); ok {
			util.WriteJSON(w, customErr.Code, dto.ErrorResponse{
				Status:  false,
				Message: customErr.Message,
				Error:   customErr.Error(),
			})
			return
		}
		util.WriteJSON(w, http.StatusInternalServerError, dto.ErrorResponse{
			Status:  false,
			Message: "Internal server error",
			Error:   err.Error(),
		})
	}

	util.WriteJSON(w, http.StatusOK, dto.MessageResponse{
		Status:  true,
		Message: "Registration successful",
	})
}

func (c *UserController) Login(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	request := dto.LoginRequest{}
	if err := json.NewDecoder(r.Body).Decode(&request); err != nil {
		util.WriteJSON(w, http.StatusBadRequest, dto.ErrorResponse{
			Status:  false,
			Message: "Invalid request body",
			Error:   err.Error(),
		})
		return
	}

	tokenData, err := c.userUsecase.Login(ctx, request)
	if err != nil {
		if customErr, ok := err.(*util.CustomError); ok {
			util.WriteJSON(w, customErr.Code, dto.ErrorResponse{
				Status:  false,
				Message: customErr.Message,
				Error:   customErr.Error(),
			})
			return
		}
		util.WriteJSON(w, http.StatusInternalServerError, dto.ErrorResponse{
			Status:  false,
			Message: "Internal server error",
			Error:   err.Error(),
		})
		return
	}

	util.WriteJSON(w, http.StatusOK, dto.DataResponse[dto.TokenData]{
		Status:  true,
		Message: "Login successful",
		Data:    *tokenData,
	})
}
