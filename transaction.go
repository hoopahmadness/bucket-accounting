package main

import (
	"time"
)

const (
	PAY      = "PAY"
	DEBIT    = "DEBIT"
	DISSOLVE = "DISSOLVE"
	ACTIVATE = "ACTIVATE"
)

type Transaction struct {
	TransactionName string
	GoalName        string
	ActionType      string
	Amount          int
	Date            time.Time
}

func newTransaction(tranType string, amount int, date time.Time) Transaction {
	var out Transaction
	out.ActionType = tranType
	out.Amount = amount
	out.Date = roundDay(date)

	return out
}
