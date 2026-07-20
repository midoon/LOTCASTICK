package repository

import (
	"context"
	"errors"

	"lotcastick-backend/internal/model"

	"gorm.io/gorm"
)

type simulationRepository struct {
	db *gorm.DB
}

func NewSimulationRepository(db *gorm.DB) model.SimulationRepository {
	return &simulationRepository{
		db: db,
	}
}

func (r *simulationRepository) WithTX(tx *gorm.DB) model.SimulationRepository {
	return &simulationRepository{
		db: tx,
	}
}

func (r *simulationRepository) Store(ctx context.Context, simulation *model.Simulation) error {
	return r.db.
		WithContext(ctx).
		Create(simulation).
		Error
}

func (r *simulationRepository) FindByID(ctx context.Context, id string) (*model.Simulation, error) {
	var simulation model.Simulation

	err := r.db.
		WithContext(ctx).
		Where("id = ?", id).
		First(&simulation).
		Error

	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, nil
		}

		return nil, err
	}

	return &simulation, nil
}

func (r *simulationRepository) FindByUserID(ctx context.Context, userID string) ([]model.Simulation, error) {
	var simulations []model.Simulation

	err := r.db.
		WithContext(ctx).
		Where("user_id = ?", userID).
		Order("created_at DESC").
		Find(&simulations).
		Error

	if err != nil {
		return nil, err
	}

	return simulations, nil
}

func (r *simulationRepository) Update(ctx context.Context, simulation *model.Simulation) error {
	return r.db.
		WithContext(ctx).
		Model(&model.Simulation{}).
		Where("id = ?", simulation.ID).
		Updates(map[string]interface{}{
			"name":           simulation.Name,
			"current_equity": simulation.CurrentEquity,
			"status":         simulation.Status,
			"passed_at":      simulation.PassedAt,
			"failed_at":      simulation.FailedAt,
			"notes":          simulation.Notes,
			"updated_at":     simulation.UpdatedAt,
		}).
		Error
}

func (r *simulationRepository) Delete(ctx context.Context, id string) error {
	return r.db.
		WithContext(ctx).
		Model(&model.Simulation{}).
		Where("id = ?", id).
		Update("deleted_at", gorm.Expr("NOW()")).
		Error
}

func (r *simulationRepository) HardDelete(ctx context.Context, id string) error {
	return r.db.
		WithContext(ctx).
		Where("id = ?", id).
		Delete(&model.Simulation{}).
		Error
}
