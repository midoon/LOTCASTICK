package model

import (
	"context"
	"time"

	"github.com/google/uuid"
	"github.com/shopspring/decimal"
	"gorm.io/gorm"
)

type DrawdownType string

const (
	DrawdownStatic   DrawdownType = "STATIC"
	DrawdownDaily    DrawdownType = "DAILY"
	DrawdownTrailing DrawdownType = "TRAILING"
	DrawdownCombined DrawdownType = "COMBINED"
)

type SimulationRule struct {
	ID                      string           `gorm:"column:id;primaryKey" json:"id"`
	SimulationID            string           `gorm:"column:simulation_id" json:"simulation_id"`
	DrawdownType            DrawdownType     `gorm:"column:drawdown_type" json:"drawdown_type"`
	MaxDrawdownPct          *decimal.Decimal `gorm:"column:max_drawdown_pct" json:"max_drawdown_pct"`
	MaxDrawdownAmount       *decimal.Decimal `gorm:"column:max_drawdown_amount" json:"max_drawdown_amount"`
	DailyDrawdownPct        *decimal.Decimal `gorm:"column:daily_drawdown_pct" json:"daily_drawdown_pct"`
	DailyDrawdownAmount     *decimal.Decimal `gorm:"column:daily_drawdown_amount" json:"daily_drawdown_amount"`
	TrailingDrawdownPct     *decimal.Decimal `gorm:"column:trailing_drawdown_pct" json:"trailing_drawdown_pct"`
	TrailingDrawdownAmount  *decimal.Decimal `gorm:"column:trailing_drawdown_amount" json:"trailing_drawdown_amount"`
	TrailingHighWaterMark   *decimal.Decimal `gorm:"column:trailing_high_water_mark" json:"trailing_high_water_mark"`
	ProfitTargetPct         decimal.Decimal  `gorm:"column:profit_target_pct" json:"profit_target_pct"`
	ProfitTargetAmount      decimal.Decimal  `gorm:"column:profit_target_amount" json:"profit_target_amount"`
	MinTradingDays          int16            `gorm:"column:min_trading_days" json:"min_trading_days"`
	ConsistencyRuleEnabled  bool             `gorm:"column:consistency_rule_enabled" json:"consistency_rule_enabled"`
	ConsistencyThresholdPct *decimal.Decimal `gorm:"column:consistency_threshold_pct" json:"consistency_threshold_pct"`
	DailyResetTimezone      string           `gorm:"column:daily_reset_timezone" json:"daily_reset_timezone"`
	DailyResetTime          time.Time        `gorm:"column:daily_reset_time" json:"daily_reset_time"`
	CreatedAt               time.Time        `gorm:"column:created_at" json:"created_at"`
	UpdatedAt               time.Time        `gorm:"column:updated_at" json:"updated_at"`
}

func (r *SimulationRule) TableName() string {
	return "simulation_rules"
}

func (r *SimulationRule) BeforeCreate(tx *gorm.DB) (err error) {
	if r.ID == "" {
		r.ID = uuid.New().String()
	}
	return nil
}

type SimulationRuleRepository interface {
	WithTX(tx *gorm.DB) SimulationRuleRepository
	Store(ctx context.Context, rule *SimulationRule) error
	FindBySimulationID(ctx context.Context, simulationID string) (*SimulationRule, error)
	Update(ctx context.Context, rule *SimulationRule) error
	Delete(ctx context.Context, simulationID string) error
}
