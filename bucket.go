package main

import (
	"time"
)

type Bucket struct {
	Name           string
	GoalAmount     int
	StartingAmount int
	CurrentAmount  int
	GoalDate       time.Time
	StartingDate   time.Time
	SavePerDay     int
	BankAccount    string
	SingleUse      bool
	Dissolved      bool
}

func newBucket(name string, amount int, startingFunds int, accountName string, singleUse bool) Bucket {
	var output Bucket
	output.Name = name
	output.GoalAmount = amount
	output.StartingAmount = startingFunds
	output.GoalDate = roundDay(time.Now()) //.Add(5)
	output.StartingDate = roundDay(time.Now())
	output.BankAccount = accountName
	output.SingleUse = singleUse
	output.updateRate()

	return output
}

func (this *Bucket) updateRate() {
	dateDiff := numDays(this.StartingDate, this.GoalDate)

	moneyDiff := this.GoalAmount - this.StartingAmount

	this.SavePerDay = moneyDiff / dateDiff

}

func (this *Bucket) transact(action Transaction) {

}

func (this *Bucket) dissolve() {
	this.Dissolved = true
}

func (this *Bucket) reactivate(input Transaction) {
	this.StartingAmount = this.CurrentAmount
	this.CurrentAmount = 0

	//to-do set dates from transaction
	//to-do reset rate

	this.Dissolved = false
}
