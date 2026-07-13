Karena request ini berasal dari JSON API, saya sarankan **DTO menggunakan `string` untuk semua nilai desimal**, bukan `decimal.Decimal`. Nanti di layer usecase baru di-parse menjadi `decimal.Decimal`.

Contohnya:

```go
package dto

type CreateSimulationRequest struct {
	Name         string                       `json:"name" validate:"required,max=150"`
	AccountSize  string                       `json:"account_size" validate:"required"`
	Currency     string                       `json:"currency" validate:"required,len=3"`
	StartedAt    string                       `json:"started_at" validate:"required"`
	Notes        *string                      `json:"notes"`
	TemplateID   *string                      `json:"template_id"`
	Rules        CreateSimulationRulesRequest `json:"rules" validate:"required"`
}

type CreateSimulationRulesRequest struct {
	DrawdownType               string  `json:"drawdown_type" validate:"required,oneof=STATIC DAILY TRAILING COMBINED"`
	MaxDrawdownPct             *string `json:"max_drawdown_pct"`
	DailyDrawdownPct           *string `json:"daily_drawdown_pct"`
	TrailingDrawdownPct        *string `json:"trailing_drawdown_pct"`
	ProfitTargetPct            string  `json:"profit_target_pct" validate:"required"`
	MinTradingDays             int16   `json:"min_trading_days"`
	ConsistencyRuleEnabled     bool    `json:"consistency_rule_enabled"`
	ConsistencyThresholdPct    *string `json:"consistency_threshold_pct"`
	DailyResetTimezone         string  `json:"daily_reset_timezone" validate:"required"`
	DailyResetTime             string  `json:"daily_reset_time" validate:"required"`
}
```

### Kenapa memakai `string` untuk decimal?

Karena payload JSON seperti ini:

```json
{
  "account_size": "100000.00",
  "profit_target_pct": "10.0000"
}
```

langsung bisa di-bind tanpa custom unmarshaler.

Di usecase tinggal:

```go
accountSize, err := decimal.NewFromString(req.AccountSize)
if err != nil {
    return err
}

profitTargetPct, err := decimal.NewFromString(req.Rules.ProfitTargetPct)
if err != nil {
    return err
}
```

---

## Kalau memakai `decimal.Decimal` langsung?

Bisa juga:

```go
AccountSize decimal.Decimal `json:"account_size"`
```

tetapi client harus mengirim:

```json
{
  "account_size": 100000.0
}
```

bukan

```json
{
  "account_size": "100000.00"
}
```

Karena contoh API kamu sudah menggunakan **string**, saya menyarankan DTO tetap menggunakan `string`, lalu konversi ke `decimal.Decimal` di layer use case. Ini juga membuat validasi dan penanganan error parsing menjadi lebih jelas.
