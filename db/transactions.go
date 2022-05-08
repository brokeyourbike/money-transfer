package db

import "gorm.io/gorm"

type TransactionsRepo interface {
}

type transactionsRepo struct {
	db *gorm.DB
}

func NewTransactionsRepo(db *gorm.DB) *transactionsRepo {
	return &transactionsRepo{db: db}
}
