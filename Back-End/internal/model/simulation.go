package model

import (
	"time"

	"github.com/google/uuid"
	"github.com/shopspring/decimal"
	"gorm.io/gorm"
)

type SimulationStatus string

const (
	SimulationStatusActive   SimulationStatus = "ACTIVE"
	SimulationStatusPassed   SimulationStatus = "PASSED"
	SimulationStatusFailed   SimulationStatus = "FAILED"
	SimulationStatusArchived SimulationStatus = "PAUSED"
)

type Simulation struct {
	ID            string           `gorm:"column:id;primaryKey" json:"id"`
	UserID        string           `gorm:"column:user_id" json:"user_id"`
	TemplateID    *string          `gorm:"column:template_id" json:"template_id"`
	Name          string           `gorm:"column:name" json:"name"`
	AccountSize   decimal.Decimal  `gorm:"column:account_size" json:"account_size"`
	CurrentEquity decimal.Decimal  `gorm:"column:current_equity" json:"current_equity"`
	Currency      string           `gorm:"column:currency" json:"currency"`
	Status        SimulationStatus `gorm:"column:status" json:"status"`
	StartedAt     time.Time        `gorm:"column:started_at" json:"started_at"`
	PassedAt      *time.Time       `gorm:"column:passed_at" json:"passed_at"`
	FailedAt      *time.Time       `gorm:"column:failed_at" json:"failed_at"`
	Notes         *string          `gorm:"column:notes" json:"notes"`
	CreatedAt     time.Time        `gorm:"column:created_at" json:"created_at"`
	UpdatedAt     time.Time        `gorm:"column:updated_at" json:"updated_at"`
	DeletedAt     *time.Time       `gorm:"column:deleted_at" json:"-"`
}

func (s *Simulation) TableName() string {
	return "simulations"
}

func (s *Simulation) BeforeCreate(tx *gorm.DB) (err error) {
	s.ID = uuid.New().String()
	return nil
}
