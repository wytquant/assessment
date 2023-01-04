package requests

import "github.com/lib/pq"

type ExpenseRequest struct {
	Title  string         `json:"title" binding:"required"`
	Amount float64        `json:"amount" binding:"required"`
	Note   string         `json:"note" binding:"required"`
	Tags   pq.StringArray `json:"tags" binding:"required"`
}
