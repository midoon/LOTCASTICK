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
		util.WriteJSON(w, http.StatusBadRequest, dto.MessageResponse{
			Status:  false,
			Message: "Invalid request body",
		})
	}

	if err := c.userUsecase.Register(ctx, request); err != nil {
		if customErr, ok := err.(*util.CustomError); ok {
			util.WriteJSON(w, customErr.Code, dto.MessageResponse{
				Status:  false,
				Message: customErr.Message,
			})
			return
		}
		util.WriteJSON(w, http.StatusInternalServerError, dto.MessageResponse{
			Status:  false,
			Message: "Internal server error",
		})
	}

	util.WriteJSON(w, http.StatusOK, dto.MessageResponse{
		Status:  true,
		Message: "Registration successful",
	})
}
