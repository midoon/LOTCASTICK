package dto

import "time"

// ========== Request DTOs ==========
type CreateSimulationRequest struct {
	Name        string                       `json:"name" validate:"required,max=150"`
	AccountSize string                       `json:"account_size" validate:"required"`
	Currency    string                       `json:"currency" validate:"required,len=3"`
	StartedAt   string                       `json:"started_at" validate:"required"`
	Notes       *string                      `json:"notes"`
	TemplateID  *string                      `json:"template_id"`
	Rules       CreateSimulationRulesRequest `json:"rules" validate:"required"`
}

type CreateSimulationRulesRequest struct {
	DrawdownType            string  `json:"drawdown_type" validate:"required,oneof=STATIC DAILY TRAILING COMBINED"`
	MaxDrawdownPct          *string `json:"max_drawdown_pct"`
	DailyDrawdownPct        *string `json:"daily_drawdown_pct"`
	TrailingDrawdownPct     *string `json:"trailing_drawdown_pct"`
	ProfitTargetPct         string  `json:"profit_target_pct" validate:"required"`
	MinTradingDays          int16   `json:"min_trading_days"`
	ConsistencyRuleEnabled  bool    `json:"consistency_rule_enabled"`
	ConsistencyThresholdPct *string `json:"consistency_threshold_pct"`
	DailyResetTimezone      string  `json:"daily_reset_timezone" validate:"required"`
	DailyResetTime          string  `json:"daily_reset_time" validate:"required"`
}

// ========== Response DTOs ==========
type SimulationResponse struct {
	ID                   string                       `json:"id"`
	Name                 string                       `json:"name"`
	AccountSize          string                       `json:"account_size"`
	CurrentEquity        string                       `json:"current_equity"`
	Currency             string                       `json:"currency"`
	Status               string                       `json:"status"`
	StartedAt            string                       `json:"started_at"` // menggungakan string karena tampilan di front-end "2026-07-13" tanggal saja tanpa jam
	PassedAt             *time.Time                   `json:"passed_at"`
	FailedAt             *time.Time                   `json:"failed_at"`
	TotalTrades          int                          `json:"total_trades"`
	TradingDaysCompleted int                          `json:"trading_days_completed"`
	PnlPct               string                       `json:"pnl_pct"`
	CreatedAt            time.Time                    `json:"created_at"`
	UpdatedAt            time.Time                    `json:"updated_at"`
	Notes                *string                      `json:"notes"`
	TemplateID           *string                      `json:"template_id"`
	Rules                SimulationRulesResponse      `json:"rules"`
	RiskStatus           SimulationRiskStatusResponse `json:"risk_status"`
}

type SimulationRulesResponse struct {
	DrawdownType            string  `json:"drawdown_type"`
	MaxDrawdownPct          *string `json:"max_drawdown_pct"`
	DailyDrawdownPct        *string `json:"daily_drawdown_pct"`
	TrailingDrawdownPct     *string `json:"trailing_drawdown_pct"`
	ProfitTargetPct         string  `json:"profit_target_pct"`
	MinTradingDays          int16   `json:"min_trading_days"`
	ConsistencyRuleEnabled  bool    `json:"consistency_rule_enabled"`
	ConsistencyThresholdPct *string `json:"consistency_threshold_pct"`
	DailyResetTimezone      string  `json:"daily_reset_timezone"`
	DailyResetTime          string  `json:"daily_reset_time"`
}

type SimulationRiskStatusResponse struct {
	Status                    string                        `json:"status"`
	Equity                    string                        `json:"equity"`
	DrawdownUsedPct           string                        `json:"drawdown_used_pct"`
	DrawdownRemainingPct      string                        `json:"drawdown_remaining_pct"`
	DrawdownRemainingAmount   string                        `json:"drawdown_remaining_amount"`
	DailyDrawdownUsedPct      string                        `json:"daily_drawdown_used_pct"`
	DailyDrawdownRemainingPct string                        `json:"daily_drawdown_remaining_pct"`
	TrailingFloor             string                        `json:"trailing_floor"`
	HighWaterMark             string                        `json:"high_water_mark"`
	ProfitTargetAmount        string                        `json:"profit_target_amount"`
	ProfitAchievedAmount      string                        `json:"profit_achieved_amount"`
	ProfitAchievedPct         string                        `json:"profit_achieved_pct"`
	TradingDaysCompleted      int                           `json:"trading_days_completed"`
	TradingDaysRequired       int                           `json:"trading_days_required"`
	ConsistencyScorePct       string                        `json:"consistency_score_pct"`
	Violations                []SimulationViolationResponse `json:"violations"`
}

type SimulationViolationResponse struct {
	ViolationType     string    `json:"violation_type"`
	EquityAtViolation string    `json:"equity_at_violation"`
	OccurredAt        time.Time `json:"occurred_at"`
}
