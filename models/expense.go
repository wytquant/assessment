package models

import "github.com/lib/pq"

type Expense struct {
	ID     uint `gorm:"primaryKey"`
	Title  string
	Amount float64
	Note   string
	Tags   pq.StringArray `gorm:"type:text[]"`
}

func (e *Expense) TableName() string {
	return "expenses"
}
