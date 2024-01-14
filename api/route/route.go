package route

import (
	"net/http"
	fetchuserbalance "txcrawlerdegen/api/handler/fetchuserbalancehandler"
	apilog "txcrawlerdegen/api/middleware"

	"github.com/gin-gonic/gin"
)

func Register(router *gin.Engine) {
	router.GET("/", func(ctx *gin.Context) {
		response := gin.H{
			"message": "Served from txcrawlerdegen",
			"status":  http.StatusOK,
			"code":    "200",
		}
		ctx.JSON(http.StatusOK, response)
	})

	router.GET("/balance", apilog.LogMiddleware(), fetchuserbalance.FetchUserBalance())
	router.GET("/ethTxs", apilog.LogMiddleware(), fetchuserbalance.FetchEthTxs())

	router.GET("/nft-balances", apilog.LogMiddleware(), fetchuserbalance.FetchEthTxs())

}
