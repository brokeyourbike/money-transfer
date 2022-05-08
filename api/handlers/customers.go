package handlers

import (
	"net/http"

	"github.com/goremit/money-transfer/db"
)

type Customers struct {
	transactionsRepo db.TransactionsRepo
}

func NewCustomers(t db.TransactionsRepo) *Customers {
	return &Customers{transactionsRepo: t}
}

func (c *Customers) HandleShowTransaction(w http.ResponseWriter, r *http.Request) {
	//
}

func (c *Customers) HandleCreateTransaction(w http.ResponseWriter, r *http.Request) {
	//
}

func (c *Customers) HandleFundTransaction(w http.ResponseWriter, r *http.Request) {
	//
}

func (c *Customers) HandleCancelTransaction(w http.ResponseWriter, r *http.Request) {
	//
}
