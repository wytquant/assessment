package responses

import "github.com/lib/pq"

type ExpenseResponse struct {
	ID     uint           `json:"id"`
	Title  string         `json:"title"`
	Amount float64        `json:"amount"`
	Note   string         `json:"note"`
	Tags   pq.StringArray `json:"tags"`
}
