package handler

import (
	"context"
	"net/http"
	"time"

	"financial-ledger/internal/core/model"
	"financial-ledger/internal/core/service"
	"financial-ledger/internal/dto"

	"github.com/gin-gonic/gin"
)

func CreateTransaction(c *gin.Context) {
	var req dto.CreateTransactionRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid request", "details": err.Error()})
		return
	}

	tx := model.Transaction{
		ID:        "",
		AccountID: req.AccountID,
		Type:      req.Type,
		Amount:    req.Amount,
		CreatedAt: time.Now().UTC(),
	}

	ctx, cancel := context.WithTimeout(c.Request.Context(), 5*time.Second)
	defer cancel()

	created, err := service.CreateTransaction(ctx, tx)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "insert failed", "details": err.Error()})
		return
	}
	c.JSON(http.StatusCreated, created)
}

func GetTransactions(c *gin.Context) {
	var q dto.TransactionQuery
	if err := c.ShouldBindQuery(&q); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid query params", "details": err.Error()})
		return
	}

	var fromT, toT time.Time
	var fromPtr, toPtr *time.Time
	var err error
	if q.From != "" {
		fromT, err = time.Parse(time.RFC3339, q.From)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": "invalid from param", "details": err.Error()})
			return
		}
		fromPtr = &fromT
	}
	if q.To != "" {
		toT, err = time.Parse(time.RFC3339, q.To)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": "invalid to param", "details": err.Error()})
			return
		}
		toPtr = &toT
	}

	ctx, cancel := context.WithTimeout(c.Request.Context(), 5*time.Second)
	defer cancel()

	txs, err := service.ListTransactions(ctx, q.AccountID, fromPtr, toPtr)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "query failed", "details": err.Error()})
		return
	}
	c.JSON(http.StatusOK, txs)
}
