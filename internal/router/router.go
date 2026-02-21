package router

import (
	"fmt"
	"net/http"
	"os"
	"time"

	"financial-ledger/internal/handler"

	"github.com/gin-gonic/gin"
)

func Router() *http.Server {
	router := gin.Default()

	api := router.Group("/api")
	{
		v1 := api.Group("/v1")
		{
			accounts := v1.Group("/accounts")
			{
				accounts.POST("/", handler.CreateAccount)
				accounts.GET("/", handler.GetAccounts)
				accounts.GET("/:accountId", handler.GetAccount)
				accounts.GET("/:accountId/balance", handler.GetBalance)
				accounts.GET("/:accountId/statement", handler.GetStatement)
			}

			transactions := v1.Group("/transactions")
			{
				transactions.POST("/", handler.CreateTransaction)
				transactions.GET("/", handler.GetTransactions)
			}
		}
	}

	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}

	return &http.Server{
		Addr:              fmt.Sprintf(":%s", port),
		Handler:           router,
		ReadHeaderTimeout: 5 * time.Second,
	}
}
