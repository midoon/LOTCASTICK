package controller

import (
	"encoding/json"
	"lotcastick-backend/internal/dto"
	"lotcastick-backend/internal/model"
	"lotcastick-backend/internal/util"
	"net/http"
)

type SimulationController struct {
	simulationUsecase model.SimulationUsecase
}

func NewSimulationController(simulationUsecase model.SimulationUsecase) *SimulationController {
	return &SimulationController{simulationUsecase: simulationUsecase}
}

func (c *SimulationController) CreateSimulation(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	userID := r.Context().Value("userID").(string)
	request := dto.CreateSimulationRequest{}

	if err := json.NewDecoder(r.Body).Decode(&request); err != nil {
		util.WriteJSON(w, http.StatusBadRequest, dto.ErrorResponse{
			Status:  false,
			Message: "Invalid request body",
			Error:   err.Error(),
		})
		return
	}

	if data, err := c.simulationUsecase.CreateSimulation(ctx, request, userID); err != nil {
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
	} else {
		// nanti kita ubah response nya ke dto.SimulationResponse
		response := dto.SimulationCreateResponse{
			SimulationID: data.SimulationID,
		}
		util.WriteJSON(w, http.StatusOK, dto.DataResponse[dto.SimulationCreateResponse]{
			Status:  true,
			Message: "Simulation created successfully",
			Data:    response,
		})
	}
}
