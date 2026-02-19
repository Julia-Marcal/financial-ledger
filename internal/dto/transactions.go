package dto

type TransactionQuery struct {
	From string `form:"from" binding:"omitempty"`
	To   string `form:"to" binding:"omitempty"`
}
