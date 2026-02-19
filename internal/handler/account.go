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

func CreateAccount(c *gin.Context) {
	var req model.Account
	if err := c.BindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid request"})
		return
	}

	ctx, cancel := context.WithTimeout(c.Request.Context(), 5*time.Second)
	defer cancel()

	created, err := service.CreateAccount(ctx, req)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "insert failed"})
		return
	}

	c.JSON(http.StatusCreated, created)
}

func GetAccount(c *gin.Context) {
	var uri dto.AccountURI
	if err := c.ShouldBindUri(&uri); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid account id", "details": err.Error()})
		return
	}

	ctx, cancel := context.WithTimeout(c.Request.Context(), 5*time.Second)
	defer cancel()

	acc, err := service.GetAccount(ctx, uri.AccountID)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "account not found"})
		return
	}
	c.JSON(http.StatusOK, acc)
}

func GetAccounts(c *gin.Context) {
	ctx, cancel := context.WithTimeout(c.Request.Context(), 5*time.Second)
	defer cancel()

	accounts, err := service.ListAccounts(ctx)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "query failed"})
		return
	}
	c.JSON(http.StatusOK, accounts)
}

func GetBalance(c *gin.Context) {
	var uri dto.AccountURI
	if err := c.ShouldBindUri(&uri); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid account id", "details": err.Error()})
		return
	}

	ctx, cancel := context.WithTimeout(c.Request.Context(), 10*time.Second)
	defer cancel()

	sum, err := service.GetBalance(ctx, uri.AccountID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "query failed"})
		return
	}
	c.JSON(http.StatusOK, gin.H{"accountId": uri.AccountID, "balance": sum})
}

func GetStatement(c *gin.Context) {
	var uri dto.AccountURI
	if err := c.ShouldBindUri(&uri); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid account id", "details": err.Error()})
		return
	}

	var q dto.StatementQuery
	if err := c.ShouldBindQuery(&q); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid query params", "details": err.Error()})
		return
	}

	var fromT, toT time.Time
	var err error
	if q.From != "" {
		fromT, err = time.Parse(time.RFC3339, q.From)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": "invalid from param", "details": err.Error()})
			return
		}
	}
	if q.To != "" {
		toT, err = time.Parse(time.RFC3339, q.To)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": "invalid to param", "details": err.Error()})
			return
		}
	}

	ctx, cancel := context.WithTimeout(c.Request.Context(), 10*time.Second)
	defer cancel()
	var fromPtr, toPtr *time.Time
	if !fromT.IsZero() {
		fromPtr = &fromT
	}
	if !toT.IsZero() {
		toPtr = &toT
	}

	txs, balance, err := service.GetStatement(ctx, uri.AccountID, fromPtr, toPtr)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "query failed"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"accountId": uri.AccountID, "transactions": txs, "balance": balance})
}
