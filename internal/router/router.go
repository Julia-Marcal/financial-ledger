package router

import (
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
)

func Router() *http.Server {
	router := gin.Default()

	api := router.Group("/api")
	{
		authorized := api.Group("/v1/").Use()
		{
			authorized.GET("/", func(c *gin.Context) {
				time.Sleep(1 * time.Second)
				c.String(http.StatusOK, "Welcome Gin Server")
			})
		}
	}

	return &http.Server{
		Addr:              ":8080",
		Handler:           router,
		ReadHeaderTimeout: 5 * time.Second,
	}
}
