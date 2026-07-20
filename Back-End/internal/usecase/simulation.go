package usecase

import (
	"context"
	"fmt"
	"lotcastick-backend/internal/database"
	"lotcastick-backend/internal/dto"
	"lotcastick-backend/internal/model"
	"lotcastick-backend/internal/util"
	"net/http"
	"time"

	"github.com/go-playground/validator/v10"
	"github.com/google/uuid"
	"github.com/shopspring/decimal"
	"github.com/spf13/viper"
	"gorm.io/gorm"
)

type simulationUsecase struct {
	tx                 database.Transaction
	simulationRepo     model.SimulationRepository
	simulationRuleRepo model.SimulationRuleRepository
	validate           *validator.Validate
	viperConfig        *viper.Viper
}

func NewSimulationUsecase(tx database.Transaction, simulationRepo model.SimulationRepository, simulationRuleRepo model.SimulationRuleRepository, validate *validator.Validate, viperConfig *viper.Viper) model.SimulationUsecase {
	return &simulationUsecase{
		tx:                 tx,
		simulationRepo:     simulationRepo,
		simulationRuleRepo: simulationRuleRepo,
		validate:           validate,
		viperConfig:        viperConfig,
	}
}

// CreateSimulation implements [model.SimulationUsecase].
func (s *simulationUsecase) CreateSimulation(ctx context.Context, req dto.CreateSimulationRequest, userID string) (*dto.SimulationCreateResponse, error) {
	if err := s.validate.Struct(req); err != nil {
		return nil, util.NewCustomError(http.StatusBadRequest, "validation error", err)
	}

	accountSize, err := decimal.NewFromString(req.AccountSize)
	if err != nil {
		return nil, util.NewCustomError(http.StatusBadRequest, "invalid account size", err)
	}

	startedAt, err := time.Parse("2006-01-02", req.StartedAt)
	if err != nil {
		return nil, util.NewCustomError(http.StatusBadRequest, "invalid started at date", err)
	}
	// buat ID untuk simulation secara manual

	simulationID := uuid.New().String()
	simulation := &model.Simulation{
		ID:            simulationID,
		UserID:        userID,
		Name:          req.Name,
		AccountSize:   accountSize,
		CurrentEquity: accountSize,
		Currency:      req.Currency,
		Status:        model.SimulationStatusActive,
		StartedAt:     startedAt,
	}

	if req.TemplateID != nil {
		fmt.Println("Template ID:", *req.TemplateID) // Debugging line to print the template ID
		templateId := *req.TemplateID
		simulation.TemplateID = &templateId
	}

	maxDrawdownPctValue := ""
	if req.Rules.MaxDrawdownPct != nil {
		maxDrawdownPctValue = *req.Rules.MaxDrawdownPct
	}

	maxDrawdownPct, err := decimal.NewFromString(maxDrawdownPctValue)
	if err != nil {
		return nil, util.NewCustomError(http.StatusBadRequest, "invalid max drawdown percentage", err)
	}

	dailyDrawdownPctValue := ""
	if req.Rules.DailyDrawdownPct != nil {
		dailyDrawdownPctValue = *req.Rules.DailyDrawdownPct
	}

	dailyDrawdownPct, err := decimal.NewFromString(dailyDrawdownPctValue)
	if err != nil {
		return nil, util.NewCustomError(http.StatusBadRequest, "invalid daily drawdown percentage", err)
	}

	trailingDrawdownPctValue := ""
	if req.Rules.TrailingDrawdownPct != nil {
		trailingDrawdownPctValue = *req.Rules.TrailingDrawdownPct
	}

	trailingDrawdownPct, err := decimal.NewFromString(trailingDrawdownPctValue)
	if err != nil {
		return nil, util.NewCustomError(http.StatusBadRequest, "invalid trailing drawdown percentage", err)
	}

	profitTargetPct, err := decimal.NewFromString(req.Rules.ProfitTargetPct)
	if err != nil {
		return nil, util.NewCustomError(http.StatusBadRequest, "invalid profit target percentage", err)
	}

	dailyResetTime, err := time.Parse("15:04:00", req.Rules.DailyResetTime)
	if err != nil {
		return nil, util.NewCustomError(http.StatusBadRequest, "invalid daily reset time", err)
	}

	consistencyThresholdPct := decimal.Zero
	if req.Rules.ConsistencyRuleEnabled && req.Rules.ConsistencyThresholdPct != nil {
		consistencyThresholdPct, err = decimal.NewFromString(*req.Rules.ConsistencyThresholdPct)
		if err != nil {
			return nil, util.NewCustomError(http.StatusBadRequest, "invalid consistency threshold percentage", err)
		}
	}

	simulationRule := &model.SimulationRule{
		SimulationID:            simulationID,
		DrawdownType:            model.DrawdownType(req.Rules.DrawdownType),
		MaxDrawdownPct:          &maxDrawdownPct,
		DailyDrawdownPct:        &dailyDrawdownPct,
		TrailingDrawdownPct:     &trailingDrawdownPct,
		ProfitTargetPct:         profitTargetPct,
		MinTradingDays:          req.Rules.MinTradingDays,
		ConsistencyRuleEnabled:  req.Rules.ConsistencyRuleEnabled,
		DailyResetTimezone:      req.Rules.DailyResetTimezone,
		ConsistencyThresholdPct: &consistencyThresholdPct,
		DailyResetTime:          dailyResetTime,
	}

	err = s.tx.WithinTransaction(ctx, func(tx *gorm.DB) error {
		simRepo := s.simulationRepo.WithTX(tx)
		ruleRepo := s.simulationRuleRepo.WithTX(tx)

		if err := simRepo.Store(ctx, simulation); err != nil {
			return err
		}

		if err := ruleRepo.Store(ctx, simulationRule); err != nil {
			return err
		}

		return nil
	})

	// nanti kita tambahka field yang lainnya
	return &dto.SimulationCreateResponse{
		SimulationID: simulation.ID,
	}, nil
}
