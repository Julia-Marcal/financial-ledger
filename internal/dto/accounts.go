package dto

// AccountURI holds path parameters for account routes.
type AccountURI struct {
	AccountID string `uri:"accountId" binding:"required"`
}

// StatementQuery holds optional query params for account statement endpoints.
type StatementQuery struct {
	From string `form:"from" binding:"omitempty"`
	To   string `form:"to" binding:"omitempty"`
}
