package dto

type CreateTransactionRequest struct {
	AccountID string `json:"account_id" binding:"required"`
	Type      string `json:"type" binding:"required,oneof=credit debit"`
	Amount    int64  `json:"amount" binding:"required,min=1"`
}

type TransactionQuery struct {
	AccountID string `form:"account_id" binding:"omitempty"`
	From      string `form:"from" binding:"omitempty"`
	To        string `form:"to" binding:"omitempty"`
}
