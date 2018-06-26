package main

import (
	"time"
)

type Report struct {
	Name               string
	TransactionHistory []Transaction
	BalanceHistory     []int
	CurrentDate        time.Time
}

func newReport(name string, currDate time.Time, startingAmount int) Report {
	var out Report
	out.Name = name
	out.CurrentDate = currDate
	out.BalanceHistory = []int{startingAmount}
	fakeTrans := newTransaction(ACTIVATE, startingAmount, currDate)
	out.TransactionHistory = []Transaction{fakeTrans}

	return out
}

func (this *Report) reactivate(input Transaction) {
}

func (this Report) export(input Transaction) {
	//creates a .csv file of the report
	//each export is up until a specific paycheck transaction
	//if a report needs to be recalculated, it will also re-do old exports
	//PAYCHECKDATE_EXPORTDATE.csv

}
