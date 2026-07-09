Saya melihat ada satu masalah pada desain response dashboard yang sekarang.

Response tersebut **bukan sekedar entity Dashboard**, tetapi merupakan **gabungan (aggregate/read model)** dari banyak tabel:

- simulations
- simulation_violations
- trades
- trade_images
- strategies
- daily_statistics

Artinya saya **tidak menyarankan** membuat satu struct `Dashboard` yang isinya asal copy response API. Lebih baik mengikuti prinsip **DTO per feature** sehingga nanti service, repository, dan handler tetap bersih.

Saya biasanya membaginya seperti berikut.

```
internal/
├── dto/
│   ├── common.go          // DataResponse, ErrorResponse, dll
│   ├── dashboard.go       // dashboard response
│   ├── simulation.go
│   ├── trade.go
│   ├── statistics.go
│   └── risk.go
```

Kemudian isi `dashboard.go`.

```go
package dto

import "time"

type DashboardResponse struct {
	Simulation      DashboardSimulation   `json:"simulation"`
	RiskStatus      DashboardRiskStatus   `json:"risk_status"`
	EquityCurve     []EquityCurvePoint    `json:"equity_curve"`
	DailyPnL        []DailyPnL            `json:"daily_pnl"`
	CalendarHeatmap []CalendarHeatmapDay  `json:"calendar_heatmap"`
	Statistics      DashboardStatistics   `json:"statistics"`
	RecentTrades    []RecentTrade         `json:"recent_trades"`
}
```

Lalu simulation.

```go
type DashboardSimulation struct {
	ID                     string     `json:"id"`
	Name                   string     `json:"name"`
	AccountSize            string     `json:"account_size"`
	CurrentEquity          string     `json:"current_equity"`
	Currency               string     `json:"currency"`
	Status                 string     `json:"status"`
	StartedAt              time.Time  `json:"started_at"`
	PassedAt               *time.Time `json:"passed_at,omitempty"`
	FailedAt               *time.Time `json:"failed_at,omitempty"`
	TotalTrades            int        `json:"total_trades"`
	TradingDaysCompleted   int        `json:"trading_days_completed"`
	PnlPct                 string     `json:"pnl_pct"`
	CreatedAt              time.Time  `json:"created_at"`
	UpdatedAt              time.Time  `json:"updated_at"`
}
```

Risk Status.

```go
type DashboardRiskStatus struct {
	Status                     string                 `json:"status"`
	Equity                     string                 `json:"equity"`
	DrawdownUsedPct            string                 `json:"drawdown_used_pct"`
	DrawdownRemainingPct       string                 `json:"drawdown_remaining_pct"`
	DrawdownRemainingAmount    string                 `json:"drawdown_remaining_amount"`
	DailyDrawdownUsedPct       string                 `json:"daily_drawdown_used_pct"`
	DailyDrawdownRemainingPct  string                 `json:"daily_drawdown_remaining_pct"`
	TrailingFloor              string                 `json:"trailing_floor"`
	HighWaterMark              string                 `json:"high_water_mark"`
	ProfitTargetAmount         string                 `json:"profit_target_amount"`
	ProfitAchievedAmount       string                 `json:"profit_achieved_amount"`
	ProfitAchievedPct          string                 `json:"profit_achieved_pct"`
	TradingDaysCompleted       int                    `json:"trading_days_completed"`
	TradingDaysRequired        int                    `json:"trading_days_required"`
	ConsistencyScorePct        string                 `json:"consistency_score_pct"`
	Violations                 []SimulationViolation  `json:"violations"`
}

type SimulationViolation struct {
	ViolationType     string    `json:"violation_type"`
	EquityAtViolation string    `json:"equity_at_violation"`
	OccurredAt        time.Time `json:"occurred_at"`
}
```

Equity Curve.

```go
type EquityCurvePoint struct {
	Timestamp time.Time `json:"timestamp"`
	Equity    string    `json:"equity"`
	TradeID   string    `json:"trade_id"`
}
```

Daily PnL.

```go
type DailyPnL struct {
	Date       time.Time `json:"date"`
	PnL        string    `json:"pnl"`
	TradeCount int       `json:"trade_count"`
}
```

Calendar.

```go
type CalendarHeatmapDay struct {
	Date       time.Time `json:"date"`
	PnL        string    `json:"pnl"`
	TradeCount int       `json:"trade_count"`
	WinRate    string    `json:"win_rate"`
}
```

Statistics.

```go
type DashboardStatistics struct {
	TotalTrades            int    `json:"total_trades"`
	WinningTrades          int    `json:"winning_trades"`
	LosingTrades           int    `json:"losing_trades"`
	WinRatePct             string `json:"win_rate_pct"`
	ProfitFactor           string `json:"profit_factor"`
	Expectancy             string `json:"expectancy"`
	AverageRR              string `json:"average_rr"`
	GrossProfit            string `json:"gross_profit"`
	GrossLoss              string `json:"gross_loss"`
	NetPnL                 string `json:"net_pnl"`
	BestDayPnL             string `json:"best_day_pnl"`
	WorstDayPnL            string `json:"worst_day_pnl"`
	AvgDailyPnL            string `json:"avg_daily_pnl"`
	MaxDrawdownReachedPct  string `json:"max_drawdown_reached_pct"`
	CurrentWinStreak       int    `json:"current_win_streak"`
	CurrentLossStreak      int    `json:"current_loss_streak"`
}
```

Trade Image.

```go
type TradeImage struct {
	ID            string    `json:"id"`
	URL           string    `json:"url"`
	Label         string    `json:"label"`
	MimeType      string    `json:"mime_type"`
	FileSizeBytes int64     `json:"file_size_bytes"`
	SortOrder     int       `json:"sort_order"`
	CreatedAt     time.Time `json:"created_at"`
}
```

Recent Trade.

```go
type RecentTrade struct {
	ID           string       `json:"id"`
	SimulationID string       `json:"simulation_id"`
	StrategyID   string       `json:"strategy_id"`
	StrategyName string       `json:"strategy_name"`

	Symbol       string       `json:"symbol"`
	Direction    string       `json:"direction"`

	EntryPrice   string       `json:"entry_price"`
	ExitPrice    string       `json:"exit_price"`

	StopLoss     string       `json:"stop_loss"`
	TakeProfit   string       `json:"take_profit"`

	LotSize      string       `json:"lot_size"`

	PnL          string       `json:"pnl"`
	PnLPct       string       `json:"pnl_pct"`

	RiskAmount   string       `json:"risk_amount"`
	RRRatio      string       `json:"rr_ratio"`

	Session       string      `json:"session"`

	EntryTime     time.Time   `json:"entry_time"`
	ExitTime      time.Time   `json:"exit_time"`
	TradeDate     time.Time   `json:"trade_date"`

	Notes         string       `json:"notes"`
	Tags          []string     `json:"tags"`

	Images        []TradeImage `json:"images"`

	CreatedAt     time.Time    `json:"created_at"`
	UpdatedAt     time.Time    `json:"updated_at"`
}
```

Kemudian endpoint tinggal menggunakan generic response yang sudah Anda miliki tanpa perlu membuat response baru.

```go
func (h *DashboardHandler) GetDashboard(c *fiber.Ctx) error {

    dashboard := dto.DashboardResponse{
        // fill data
    }

    return c.Status(fiber.StatusOK).JSON(dto.DataResponse[dto.DashboardResponse]{
        Status:  true,
        Message: "Dashboard data retrieved",
        Data:    dashboard,
    })
}
```

## Saran perbaikan yang lebih penting

Karena aplikasi Anda cukup besar (Trading Journal + Analytics + Prop Firm Simulator), saya justru menyarankan **memisahkan DTO berdasarkan domain**, bukan berdasarkan endpoint.

Contoh struktur yang lebih scalable:

```
internal/
├── dto/
│
├── common/
│   ├── response.go
│   └── pagination.go
│
├── simulation/
│   ├── request.go
│   ├── response.go
│   └── model.go
│
├── dashboard/
│   ├── response.go
│   ├── statistics.go
│   ├── equity_curve.go
│   ├── risk.go
│   └── calendar.go
│
├── trade/
│   ├── request.go
│   ├── response.go
│   ├── image.go
│   └── statistics.go
│
├── analytics/
│   ├── response.go
│   ├── symbol.go
│   ├── strategy.go
│   ├── session.go
│   └── drawdown.go
│
├── auth/
│   ├── request.go
│   └── response.go
```

Dengan struktur tersebut, ketika nanti Anda menambahkan fitur seperti **Analytics**, **Performance by Symbol**, **Performance by Strategy**, atau **Prop Firm Rules**, DTO tidak akan menjadi satu file berisi ratusan struct dan proyek akan jauh lebih mudah dipelihara.
