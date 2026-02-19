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
			v1.GET("/", func(c *gin.Context) {
				time.Sleep(1 * time.Second)
				c.String(http.StatusOK, "Welcome Gin Server")
			})

			accounts := v1.Group("/accounts")
			{
				accounts.POST("/", handler.CreateAccount)
				accounts.GET("/", handler.GetAccounts)
				accounts.GET("/:accountId", handler.GetAccount)
				accounts.GET("/:accountId/balance", handler.GetBalance)
				accounts.GET("/:accountId/statement", handler.GetStatement)
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
