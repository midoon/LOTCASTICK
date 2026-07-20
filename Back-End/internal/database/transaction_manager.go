package database

import (
	"context"

	"gorm.io/gorm"
)

type Transaction interface {
	WithinTransaction(
		ctx context.Context,
		fn func(tx *gorm.DB) error,
	) error
}

type transactionManager struct {
	db *gorm.DB
}

func NewTransactionManager(db *gorm.DB) Transaction {
	return &transactionManager{
		db: db,
	}
}

func (t *transactionManager) WithinTransaction(
	ctx context.Context,
	fn func(tx *gorm.DB) error,
) error {
	return t.db.WithContext(ctx).Transaction(func(tx *gorm.DB) error {
		return fn(tx)
	})
}
