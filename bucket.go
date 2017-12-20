package main

import (
	"time"
)

type Bucket struct {
	goalName           string
	goalAmount         int
	goalDate           time.Time
	currentAmountSaved int
	savePerDay         int
}

func newBucket(name string, amount int, date time.Time) Bucket {
	var newBucket Bucket
	newBucket.goalName = name
	newBucket.goalAmount = amount
	newBucket.goalDate = date

	return newBucket
}

func (this *Bucket) moveMoney(cash int) {
	this.currentAmountSaved += cash
}

func (this *Bucket) setByRate(rate int) {
	this.savePerDay = rate

}
