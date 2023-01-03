package requests

import "github.com/lib/pq"

type ExpenseRequest struct {
	Title  string         `json:"title"`
	Amount float64        `json:"amount"`
	Note   string         `json:"note"`
	Tags   pq.StringArray `json:"tags"`
}
