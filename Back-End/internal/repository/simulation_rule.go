package repository

import (
	"context"
	"errors"

	"lotcastick-backend/internal/model"

	"gorm.io/gorm"
)

type simulationRuleRepository struct {
	db *gorm.DB
}

func NewSimulationRuleRepository(db *gorm.DB) model.SimulationRuleRepository {
	return &simulationRuleRepository{
		db: db,
	}
}

func (r *simulationRuleRepository) Store(ctx context.Context, rule *model.SimulationRule) error {
	return r.db.
		WithContext(ctx).
		Create(rule).
		Error
}

func (r *simulationRuleRepository) FindBySimulationID(ctx context.Context, simulationID string) (*model.SimulationRule, error) {
	var rule model.SimulationRule

	err := r.db.
		WithContext(ctx).
		Where("simulation_id = ?", simulationID).
		First(&rule).
		Error

	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, nil
		}

		return nil, err
	}

	return &rule, nil
}

func (r *simulationRuleRepository) Update(ctx context.Context, rule *model.SimulationRule) error {
	return r.db.
		WithContext(ctx).
		Model(&model.SimulationRule{}).
		Where("id = ?", rule.ID).
		Updates(map[string]interface{}{
			"drawdown_type":             rule.DrawdownType,
			"max_drawdown_pct":          rule.MaxDrawdownPct,
			"max_drawdown_amount":       rule.MaxDrawdownAmount,
			"daily_drawdown_pct":        rule.DailyDrawdownPct,
			"daily_drawdown_amount":     rule.DailyDrawdownAmount,
			"trailing_drawdown_pct":     rule.TrailingDrawdownPct,
			"trailing_drawdown_amount":  rule.TrailingDrawdownAmount,
			"trailing_high_water_mark":  rule.TrailingHighWaterMark,
			"profit_target_pct":         rule.ProfitTargetPct,
			"profit_target_amount":      rule.ProfitTargetAmount,
			"min_trading_days":          rule.MinTradingDays,
			"consistency_rule_enabled":  rule.ConsistencyRuleEnabled,
			"consistency_threshold_pct": rule.ConsistencyThresholdPct,
			"daily_reset_timezone":      rule.DailyResetTimezone,
			"daily_reset_time":          rule.DailyResetTime,
			"updated_at":                rule.UpdatedAt,
		}).
		Error
}

func (r *simulationRuleRepository) Delete(ctx context.Context, simulationID string) error {
	return r.db.
		WithContext(ctx).
		Where("simulation_id = ?", simulationID).
		Delete(&model.SimulationRule{}).
		Error
}
